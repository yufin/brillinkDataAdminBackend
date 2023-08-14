package syncdependency

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/sftp"
	"go-admin/app/rc/models"
	"go-admin/app/rc/service/dto"
	cModels "go-admin/common/models"
	"go-admin/pkg/natsclient"
	cUtils "go-admin/utils"
	"gorm.io/gorm"
	"io"
	"path"
	"regexp"
	"sync"
	"sync/atomic"
)

const (
	UscIdPattern     string = `^[a-zA-Z0-9]{18}$`
	YearMonthPattern string = `^[0-9]{6}$`
	ConcurrencyLimit int    = 5
)

var soctRunning int32

type SyncOriginContentTask struct {
}

func (t SyncOriginContentTask) Exec(arg interface{}) error {
	if atomic.LoadInt32(&soctRunning) == 1 {
		log.Info("SyncOriginContentTask任务已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&soctRunning, 1)
	defer atomic.StoreInt32(&soctRunning, 0)

	err := SyncOriginJsonContent()
	if err != nil {
		log.Errorf("TASK SyncOriginJsonContent Failed:%s \r\n", err)
	}
	return nil
}

// SyncOriginJsonContent 同步微众企业风控数据json数据至数据库
func SyncOriginJsonContent() error {
	// 1. 获取sftp连接
	sftpClientP, err := cUtils.GetSftpClient()
	if err != nil {
		log.Errorf("GetSftpClient Failed:%s \r\n", err)
		return err
	}

	defer cUtils.CloseShhConn()
	defer sftpClientP.Close()

	// 2. 遍历sftp目录，获取所有的json所属文件文件夹路径
	dirNames, err := sftpClientP.Glob("/taxDataPreloanFile/*")
	if err != nil {
		log.Error(err)
		//panic(err)
		return err
	}

	// 筛选出符合条件的目录
	matchedDirs := FilteringDirName(dirNames, UscIdPattern)

	// 获取目录信息
	var dirInfos []DirInfo
	for _, dirPath := range matchedDirs {
		GetDirInfo(dirPath, sftpClientP, &dirInfos)
	}

	// 读取数据库，找出不存在数据库中的文件信息，存入数据库
	var wg sync.WaitGroup
	var mutex sync.Mutex
	limitCh := make(chan struct{}, ConcurrencyLimit)

	var tb models.RcOriginContent
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	for i, _ := range dirInfos {
		limitCh <- struct{}{}
		wg.Add(1)
		go func(index int, cdb *gorm.DB) {
			defer wg.Done()
			mutex.Lock()
			dirInfoP := &dirInfos[index]
			mutex.Unlock()
			err := CheckIfInfoRecorded(dirInfoP, cdb)
			if err != nil {
				log.Errorf("CheckIfInfoRecorded Error: %s \r\n", err)
			}
			<-limitCh
		}(i, db)
	}
	wg.Wait()

	// 遍历dirInfos,读取文件录入数据库
	for _, dirInfo := range dirInfos {
		if dirInfo.notExist == true {
			content := string(GetFileContentFromSftp(sftpClientP, dirInfo.DataFilePath))
			var en string
			en, err = parseEnterpriseName(content)
			if err != nil {
				log.Errorf("parseEnterpriseName Error: %s \r\n", err)
				return err
			}

			insertReq := dto.RcOriginContentInsertReq{
				UscId:          dirInfo.UscId,
				EnterpriseName: en,
				YearMonth:      dirInfo.YearMonth,
				Content:        content,
				StatusCode:     0,
				ControlBy:      cModels.ControlBy{CreateBy: 0},
			}
			var data models.RcOriginContent
			insertReq.Generate(&data)
			err := db.Model(&data).Create(&data).Error
			if err != nil {
				log.Errorf("SyncRcOriginContent Insert Error: %s \r\n", err)
			} else {
				log.Infof("SyncRcOriginContent Insert Success: USCID:%s; ImportedAt:%s \r\n", insertReq.UscId, insertReq.YearMonth)
				// publish data.Id
				err := func() error {
					msg := make([]byte, 8)
					binary.BigEndian.PutUint64(msg, uint64(data.Id))
					_, err := natsclient.TaskJs.Publish(natsclient.TopicContentNew, msg)
					if err != nil {
						return err
					}
					return nil
				}()
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func GetFileContentFromSftp(client *sftp.Client, filePath string) []byte {
	remoteFile, err := client.Open(filePath)
	if err != nil {
		log.Errorf("Sftp client Open file error!:%s \r\n", err)
		panic(err)
	}
	defer remoteFile.Close()
	fileContent, err := io.ReadAll(remoteFile)
	if err != nil {
		log.Errorf("Read file content error!:%s \r\n", err)
		panic(err)
	}
	return fileContent
}

// FilteringDirName 匹配目录名
func FilteringDirName(pathArray []string, patternExpr string) []string {
	pattern := regexp.MustCompile(patternExpr)
	var tempDirs []string
	for _, dir := range pathArray {
		if pattern.MatchString(path.Base(dir)) {
			tempDirs = append(tempDirs, dir)
		}
	}
	return tempDirs
}

type DirInfo struct {
	UscId        string
	YearMonth    string
	DataFilePath string
	notExist     bool
}

func CheckIfInfoRecorded(dirInfo *DirInfo, db *gorm.DB) error {
	var tb models.RcOriginContent
	var count int64
	err := db.Model(&tb).
		Where("usc_id = ?", dirInfo.UscId).
		Where("`year_month` = ?", dirInfo.YearMonth).
		Count(&count).
		Error
	if err != nil {
		return err
	}
	if count == 0 {
		(*dirInfo).notExist = true
	} else {
		(*dirInfo).notExist = false
	}
	return nil
}

func GetDirInfo(dirPath string, client *sftp.Client, dirInfos *[]DirInfo) {
	// open dir by dirPath, then get children dir's fileName
	uscId := path.Base(dirPath)
	dirs, err := client.Glob(dirPath + "/*")
	if err != nil {
		log.Error(err)
		panic(err)
	}
	matchedDirs := FilteringDirName(dirs, YearMonthPattern)
	for _, dir := range matchedDirs {
		// get child path of dir which ended with .json
		JsonFilePathArray, err := client.Glob(dir + "/*.json")
		if err != nil {
			log.Error(err)
			panic(err)
		}
		if len(JsonFilePathArray) == 0 {
			return
		}
		*dirInfos = append(*dirInfos, DirInfo{
			UscId:        uscId,
			YearMonth:    fmt.Sprintf("%s-%s", path.Base(dir)[:4], path.Base(dir)[4:]),
			DataFilePath: JsonFilePathArray[0],
		})
	}
}

func parseEnterpriseName(content string) (string, error) {
	var contentMap map[string]any
	err := json.Unmarshal([]byte(content), &contentMap)
	if err != nil {
		return "", err
	}
	s := contentMap["impExpEntReport"].(map[string]any)["businessInfo"].(map[string]any)["QYMC"]
	// determine enterprise name is nil or string
	if _, ok := s.(string); ok {
		return s.(string), nil
	}
	return "", fmt.Errorf("enterprise name is not string")
}

package task

import (
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/google/uuid"
	"github.com/pkg/sftp"
	"go-admin/app/rskc/service"
	"go-admin/app/rskc/service/dto"
	"go-admin/app/rskc/utils"
	cDto "go-admin/common/dto"
	"go-admin/common/models"
	"io/ioutil"
	"path"
	"regexp"
	"sync"
)

const (
	UscIdPattern     = `^[a-zA-Z0-9]{18}$`
	YearMonthPattern = `^[0-9]{6}$`
)

type SyncOriginContentTask struct {
}

func (t SyncOriginContentTask) Exec(arg interface{}) error {
	err := SyncOriginJsonContent()
	return err
}

// SyncOriginJsonContent 同步微众企业风控数据json数据至数据库
func SyncOriginJsonContent() error {
	// 1. 获取sftp连接
	sftpClientP, err := utils.GetSftpClient()
	if err != nil {
		log.Errorf("GetSftpClient Failed:%s \r\n", err)
		return err
	}
	defer utils.CloseShhConn()
	defer func(sftpClientP *sftp.Client) {
		err := sftpClientP.Close()
		if err != nil {
			log.Error(err)
			panic(err)
		}
	}(sftpClientP)

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
	for _, dirInfo := range dirInfos {
		wg.Add(1)
		go func(dirInfoP *DirInfo) {
			defer wg.Done()
			err := CheckIfInfoRecorded(dirInfoP)
			if err != nil {
				log.Errorf("CheckIfInfoRecorded Error: %s \r\n", err)
			}
		}(&dirInfo)
	}
	wg.Wait()

	// 遍历dirInfos,读取文件录入数据库
	s := service.OriginContent{}
	for _, dirInfo := range dirInfos {
		if dirInfo.notExist == true {
			insertReq := dto.OriginContentInsertReq{
				ContentId:         uuid.New().String(),
				UscId:             dirInfo.UscId,
				YearMonth:         dirInfo.YearMonth,
				OriginJsonContent: string(GetFileContentFromSftp(sftpClientP, dirInfo.DataFilePath)),
				ControlBy:         models.ControlBy{CreateBy: 0},
			}
			err := s.Insert(&insertReq)
			if err != nil {
				log.Errorf("SyncRskcOriginContent Insert Error: %s \r\n", err)
			} else {
				log.Infof("SyncRskcOriginContent Insert Success: USCID:%s; ImportedAt:%s \r\n", insertReq.UscId, insertReq.YearMonth)
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
	fileContent, err := ioutil.ReadAll(remoteFile)
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

func CheckIfInfoRecorded(dirInfo *DirInfo) error {
	s := service.OriginContent{}
	req := dto.OriginContentGetPageReq{
		Pagination: cDto.Pagination{PageIndex: 1, PageSize: 10},
		UscId:      dirInfo.UscId,
		YearMonth:  dirInfo.YearMonth,
	}
	var count int64
	err := s.CountByInfo(&req, &count)
	if err != nil {
		return err
	}
	if count == 0 {
		dirInfo.notExist = true
	} else {
		dirInfo.notExist = false
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
			UscId:     uscId,
			YearMonth: path.Base(dir),
			DataFilePath: fmt.Sprintf(
				"%s-%s", JsonFilePathArray[0][:4], JsonFilePathArray[0][4:]),
		})
	}
}

package task

import (
	"context"
	"encoding/binary"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"go-admin/app/rc/models"
	cModels "go-admin/common/models"
	"go-admin/config"
	"go-admin/pkg/gotenbergclient"
	"go-admin/pkg/minioclient"
	"go-admin/pkg/natsclient"
	"go-admin/utils"
	"net/url"
	"sync/atomic"
	"time"
)

var snapshotRunning int32

type ReportSnapshotTask struct {
}

func (t ReportSnapshotTask) Exec(arg interface{}) error {
	if atomic.LoadInt32(&snapshotRunning) == 1 {
		log.Info("ReportSnapShotTask已经在执行中，跳过本次调度")
		return nil
	}
	atomic.StoreInt32(&snapshotRunning, 1)
	defer atomic.StoreInt32(&snapshotRunning, 0)

	if err := t.pubId4Snapshot(); err != nil {
		log.Errorf("pubId4Snapshot Failed:%v \r\n", err)
		return err
	}

	gtb := gotenbergclient.NewGtbClient(config.ExtConfig.PdfConvert.Gtb.Server)
	for {
		msgs, err := natsclient.SubReportSnapshot.Fetch(1, nats.MaxWait(5*time.Second))
		if err != nil {
			if err == nats.ErrTimeout {
				return nil
			} else {
				return err
			}
		}
		for _, msg := range msgs {
			depId := int64(binary.BigEndian.Uint64(msg.Data))
			if err := t.getReportSnapshot(depId, gtb, minioclient.MinioCli); err != nil {
				log.Errorf("getReportSnapshot Failed:%v \r\n", err)
				return err
			}
			if err := msg.AckSync(); err != nil {
				return err
			}
		}

	}
}

func (t ReportSnapshotTask) pubId4Snapshot() error {
	tb := models.RcReportOss{}
	tbRdd := models.RcDependencyData{}
	db := sdk.Runtime.GetDbByKey(tb.TableName())
	depIds := make([]int64, 0)
	err := db.Raw(
		`select rdd.id as dep_id
			from rc_dependency_data rdd
					 left join rc_report_oss rro on rdd.id = rro.dep_id
			where rro.id is null and rdd.deleted_at is null and rro.deleted_at is null and rdd.status_code = 0;`).
		Pluck("dep_id", &depIds).
		Error
	if err != nil {
		return err
	}
	for _, depId := range depIds {
		msg := make([]byte, 8)
		binary.BigEndian.PutUint64(msg, uint64(depId))
		_, err := natsclient.TaskJs.Publish(natsclient.TopicReportSnapshot, msg)
		if err != nil {
			return err
		}
		// update statusCode = 1
		err = db.Model(&tbRdd).Where("id = ?", depId).Update("status_code", 1).Error
		if err != nil {
			return err
		}

	}
	return nil
}

func (t ReportSnapshotTask) getReportSnapshot(depId int64, gtb gotenbergclient.GtbCli, mc minioclient.McInterface) error {
	//http://192.168.44.150:1024/login?redirect=/CrawlReport?depId=468225900752678038&lang=zh&u=admin&p=1234
	raw := fmt.Sprintf("%s%s?depId=%d&lang=%s&u=%s&p=%s",
		config.ExtConfig.PdfConvert.Report.Server,
		config.ExtConfig.PdfConvert.Report.Path,
		depId,
		"zh",
		config.ExtConfig.PdfConvert.Report.Username,
		config.ExtConfig.PdfConvert.Report.Password)
	u, err := url.Parse(raw)
	if err != nil {
		return err
	}
	resp, err := gtb.ChromiumConvert(u.String())
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return err
	}

	bucketName := config.ExtConfig.PdfConvert.Report.OssBucketName
	randomId, _ := uuid.NewRandom()
	objName := fmt.Sprintf("%d-%s.pdf", depId, randomId.String())
	err = mc.UploadFile(context.Background(), bucketName, objName, resp.Body)
	if err != nil {
		return err
	}

	// add metadata record
	ossRec := models.OssMetadata{
		Model:      cModels.Model{Id: utils.NewFlakeId()},
		ObjName:    objName,
		BucketName: bucketName,
		Endpoint:   mc.GetCli().EndpointURL().String(),
		App:        1,
	}
	db := sdk.Runtime.GetDbByKey(ossRec.TableName())
	err = db.Create(&ossRec).Error
	if err != nil {
		return err
	}
	// add report oss record.
	repRec := models.RcReportOss{
		Model:   cModels.Model{Id: utils.NewFlakeId()},
		DepId:   depId,
		OssId:   ossRec.Id,
		Version: 2,
	}
	dbDep := sdk.Runtime.GetDbByKey(repRec.TableName())
	err = dbDep.Create(&repRec).Error
	if err != nil {
		return err
	}
	return nil
}

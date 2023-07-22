package rdm

import (
	"bytes"
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"go-admin/config"
	"go-admin/utils"
	"io"
	"net/http"
	"strconv"
)

type PySidecarAhpRdm struct {
	depId int64
}

// Pipeline 流程
func (t PySidecarAhpRdm) Pipeline() error {
	f, err := t.getFactors()
	if err != nil {
		return err
	}

	res, err := t.getRes(f)
	if err != nil {
		return err
	}

	for _, level := range []int{1, 2, 3, 4} {
		resKey := "l" + strconv.Itoa(level)
		l1, dt, _, err := jsonparser.Get(res, resKey)
		if dt != jsonparser.Object {
			return errors.New("l1 is not object")
		}
		m := make(map[string]float64)
		err = json.Unmarshal(l1, &m)
		if err != nil {
			return err
		}
		if err := t.saveResult(m, level); err != nil {
			return err
		}
	}

	// L1 process

	return nil
}

func (t PySidecarAhpRdm) getRes(factor map[string]any) ([]byte, error) {
	// http post request
	u := config.ExtConfig.PySidecar.Uri + config.ExtConfig.PySidecar.AhpPath
	payload, err := json.Marshal(factor)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(u, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("http status code is not 200")
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t PySidecarAhpRdm) saveResult(factor map[string]float64, level int) error {
	tb := models.RcRdmResult{}
	db := sdk.Runtime.GetDbByKey(tb.TableName())

	for k, v := range factor {
		data := models.RcRdmResult{
			DepId: t.depId,
			Field: k,
			Level: level,
			Score: v,
		}
		data.Id = utils.NewFlakeId()

		if err := db.Create(&data).Error; err != nil {
			return err
		}
	}
	return nil
}

// getFactors 获取决策因子
func (t PySidecarAhpRdm) getFactors() (map[string]any, error) {
	dt := models.RcDependencyData{}
	db := sdk.Runtime.GetDbByKey(dt.TableName())
	if err := db.Model(&dt).First(&dt, t.depId).Error; err != nil {
		return nil, err
	}
	dtParam := models.RcDecisionParam{}
	dbParam := sdk.Runtime.GetDbByKey(dtParam.TableName())
	if err := dbParam.Model(&dtParam).
		Where("content_id = ?", dt.ContentId).
		Order("created_at desc").
		First(&dtParam).
		Error; err != nil {
		return nil, err
	}

	dtParam.LhQylx = dt.LhQylx
	dtParam.LhCylwz = dt.LhCylwz
	dtParam.MdQybq = dt.LhQybq
	dtParam.GsGdct = dt.LhGdct
	dtParam.ZxYhsxqk = dt.LhYhsx
	dtParam.ZxDsfsxqk = dt.LhSfsx
	// dtParam to map[string]any
	b, _ := json.Marshal(dtParam)
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	// remove unnecessary keys
	delete(m, "id")
	delete(m, "createBy")
	delete(m, "updateBy")
	delete(m, "createdAt")
	delete(m, "updatedAt")
	delete(m, "nsrsbh")

	return m, nil
}

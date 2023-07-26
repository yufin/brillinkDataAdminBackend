package syncdependency

import (
	"github.com/buger/jsonparser"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go-admin/app/rc/models"
	"go-admin/app/rc/service/dto"
	cModels "go-admin/common/models"
	"gorm.io/gorm"
	"strconv"
)

type decisionParamSyncProcess struct {
}

func (t decisionParamSyncProcess) Process(contentId int64) error {
	var dataContent models.RcOriginContent
	dbContent := sdk.Runtime.GetDbByKey(dataContent.TableName())
	err := dbContent.Model(&dataContent).First(&dataContent, contentId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	indexMap := make(map[string]any)
	contentBytes := []byte(dataContent.Content)
	offset, err := jsonparser.ArrayEach(contentBytes, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		indexValue, _ := jsonparser.GetString(value, "INDEX_VALUE")
		indexCode, _ := jsonparser.GetString(value, "INDEX_CODE")
		indexMap[indexCode] = indexValue
	}, "modelIndexes")
	if err != nil {
		return errors.Errorf("jsonparser error at offset=%d,contentId=%d, error,: %v", offset, contentId, err)
	}
	insertReq := dto.RcDecisionParamInsertReq{
		ContentId:               contentId,
		SwSdsnbCyrs:             t.safeGetIntFromMap(indexMap, "sw_sdsnb_cyrs"),
		GsGdct:                  t.safeGetIntFromMap(indexMap, "gs_gdct"),
		GsGdwdx:                 t.safeGetIntFromMap(indexMap, "gs_gdwdx"),
		GsFrwdx:                 t.safeGetIntFromMap(indexMap, "gs_frwdx"),
		LhCylwz:                 t.safeGetIntFromMap(indexMap, "lh_cylwz"),
		LhMdPpjzl:               t.safeGetIntFromMap(indexMap, "lh_md_ppjzl"),
		MdQybq:                  t.safeGetIntFromMap(indexMap, "md_qybq"),
		SwCwbbYyzjzzts:          t.safeGetIntFromMap(indexMap, "sw_cwbb_yyzjzzts"),
		SfFhSsqkQy:              t.safeGetIntFromMap(indexMap, "sf_fh_ssqk_qy"),
		SwJcxxNsrxypj:           t.safeGetStringFromMap(indexMap, "sw_jcxx_nsrxypj"),
		ZxYhsxqk:                t.safeGetIntFromMap(indexMap, "zx_yhsxqk"),
		ZxDsfsxqk:               t.safeGetIntFromMap(indexMap, "zx_dsfsxqk"),
		LhQylx:                  t.safeGetIntFromMap(indexMap, "lh_qylx"),
		Nsrsbh:                  &dataContent.UscId,
		SwSbNszeZzsqysds12m:     t.safeGetDecimalFromMap(indexMap, "sw_sb_nsze_zzsqysds_12m"),
		SwSbNszezzlZzsqysds12mA: t.safeGetDecimalFromMap(indexMap, "sw_sb_nszezzl_zzsqysds_12m_a"),
		SwSdsnbGzxjzzjezzl:      t.safeGetDecimalFromMap(indexMap, "sw_sdsnb_gzxjzzjezzl"),
		SwSbzsSflhypld12m:       t.safeGetDecimalFromMap(indexMap, "sw_sbzs_sflhypld_12m"),
		SwSdsnbYjfy:             t.safeGetDecimalFromMap(indexMap, "sw_sdsnb_yjfy"),
		FpJxLxfy12m:             t.safeGetDecimalFromMap(indexMap, "fp_jx_lxfy_12m"),
		SwCwbbSszb:              t.safeGetDecimalFromMap(indexMap, "sw_cwbb_sszb"),
		FpJySychjeZb12mLh:       t.safeGetDecimalFromMap(indexMap, "fp_jy_sychje_zb_12m_lh"),
		FpJxZyjyjezb12mLh:       t.safeGetDecimalFromMap(indexMap, "fp_jx_zyjyjezb_12m_lh"),
		FpXxXychjeZb12mLh:       t.safeGetDecimalFromMap(indexMap, "fp_xx_xychje_zb_12m_lh"),
		FpXxZyjyjezb12mLh:       t.safeGetDecimalFromMap(indexMap, "fp_xx_zyjyjezb_12m_lh"),
		SwSbQbxse12m:            t.safeGetDecimalFromMap(indexMap, "sw_sb_qbxse_12m"),
		SwSbQbxsezzl12m:         t.safeGetDecimalFromMap(indexMap, "sw_sb_qbxsezzl_12m"),
		SwSbLsxs12m:             t.safeGetDecimalFromMap(indexMap, "sw_sb_lsxs_12m"),
		SwCwbbChzztsCb:          t.safeGetDecimalFromMap(indexMap, "sw_cwbb_chzzts_cb"),
		SwCwbbZcfzl:             t.safeGetDecimalFromMap(indexMap, "sw_cwbb_zcfzl"),
		SwCwbbMlrzzlv:           t.safeGetDecimalFromMap(indexMap, "sw_cwbb_mlrzzlv"),
		SwCwbbJlrzzlv:           t.safeGetDecimalFromMap(indexMap, "sw_cwbb_jlrzzlv"),
		SwCwbbJzcszlv:           t.safeGetDecimalFromMap(indexMap, "sw_cwbb_jzcszlv"),
		SwJcxxClnx:              t.safeGetDecimalFromMap(indexMap, "sw_jcxx_clnx"),
		StatusCode:              0,
		ControlBy:               cModels.ControlBy{},
	}

	var dtParam models.RcDecisionParam
	insertReq.Generate(&dtParam)
	dbPram := sdk.Runtime.GetDbByKey(dtParam.TableName())
	err = dbPram.Model(&dtParam).Create(&dtParam).Error
	if err != nil {
		return err
	}
	return nil
}

func (t decisionParamSyncProcess) safeGetDecimalFromMap(m map[string]any, key string) decimal.NullDecimal {
	nullDecimal := decimal.NullDecimal{Valid: false}
	if v, ok := m[key]; ok {
		if v == nil {
			return nullDecimal
		}
		dec, err := decimal.NewFromString(v.(string))
		if err != nil {
			return nullDecimal
		}
		return decimal.NullDecimal{
			Decimal: dec,
			Valid:   true,
		}
	}
	return nullDecimal
}

func (t decisionParamSyncProcess) safeGetStringFromMap(m map[string]any, key string) *string {
	if v, ok := m[key]; ok {
		if v == nil {
			return nil
		}
		v := v.(string)
		return &v
	}
	return nil
}

func (t decisionParamSyncProcess) safeGetIntFromMap(m map[string]any, key string) *int {
	if v, ok := m[key]; ok {
		if v == nil {
			return nil
		}
		r, err := strconv.Atoi(v.(string))
		if err != nil {
			return nil
		}
		return &r
	}
	return nil
}

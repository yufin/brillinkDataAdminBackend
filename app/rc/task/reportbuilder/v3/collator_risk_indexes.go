package v3

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
	"math"
	"strconv"
)

type ClaRiskIndexes struct {
	content   *[]byte
	contentId int64
}

func (s *ClaRiskIndexes) SetContent(content *[]byte, contentId int64) {
	s.content = content
	s.contentId = contentId
}

func (s *ClaRiskIndexes) Collating() error {
	paths := []string{"impExpEntReport", "riskIndexes"}
	var idx int
	var errCb, err error
	c := *s.content

	idx = 0
	_, err = jsonparser.ArrayEach(c,
		func(value []byte, dt jsonparser.ValueType, offset int, err error) {
			if dt != jsonparser.Object {
				errCb = errors.New("value is not object")
				return
			}

			idxStr := fmt.Sprintf("[%d]", idx)
			var indexDecVal, indexValue string
			indexDecVal, errCb = jsonparser.GetString(value, "INDEX_DEC")
			if errCb != nil {
				return
			}
			indexValue, errCb = jsonparser.GetString(value, "INDEX_VALUE")
			if errCb != nil {
				return
			}

			var resetVal string
			switch indexDecVal {
			case "历史变更-企业名称变更":
				resetVal = s.lambdaReformatIndexValue("变更企业名称%s次", indexValue)
			case "历史变更-地址变更":
				resetVal = s.lambdaReformatIndexValue("变更地址%s次", indexValue)
			case "变更的风险提示\n（减资）":
				if v, err := strconv.ParseInt(indexValue, 10, 32); err != nil {
					if v > 0 {
						resetVal = s.lambdaReformatIndexValue("增资注册资本%s万", indexValue)
					} else if v < 0 {
						resetVal = s.lambdaReformatIndexValue("减资注册资本%s万", indexValue)
					}
				} else {
					resetVal = "无"
				}
			case "金融欠款纠纷":
				resetVal = s.lambdaReformatIndexValue("近5年内企业和法人作为被告%s次", indexValue)
			case "作为被告裁判文书涉案金额":
				resetVal = s.lambdaReformatIndexValue("%s（近5年企业民事裁判文书中作为被告涉案金额加总+近5年法人民事裁判文书中作为被告涉案金额加总）/（上年）营业收入", indexValue)
			case "作为原告裁判文书涉案金额":
				resetVal = s.lambdaReformatIndexValue("%s（近5年企业民事裁判文书中作为原告涉案金额加总+近5年法人民事裁判文书中作为原告涉案金额加总）/（上年）营业收入", indexValue)
			case "法人历史失信记录":
				resetVal = s.lambdaReformatIndexValue("近5年法人失信%s次", indexValue)
			case "历史被执行人记录":
				resetVal = s.lambdaReformatIndexValue("近5年企业被执行%s次", indexValue)
			case "工商处罚记录":
				resetVal = s.lambdaReformatIndexValue("近5年企业工商处罚%s次", indexValue)
			case "收入成长率":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "毛利润成长率":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "供应商评价":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "供应商稳定性":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "客户评价":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "客户稳定性":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "主营业务专注度":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "净利与毛利波动差异":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "现金比率":
				resetVal = s.lambdaReformatIndexValue("%s%%", indexValue)
			case "应收运营周转天数":
				resetVal = s.round(indexValue)
			case "应付运营周转天数":
				resetVal = s.round(indexValue)
			case "存货周转天数":
				resetVal = s.round(indexValue)
			}

			if resetVal != "" {
				toSet := fmt.Sprintf(`"%s"`, resetVal)
				*s.content, errCb = jsonparser.Set(*s.content, []byte(toSet), append(paths, idxStr, "INDEX_VALUE")...)
				if errCb != nil {
					return
				}
			}
			idx++
		}, paths...)
	if errCb != nil {
		return errCb
	}
	if err != nil {
		return err
	}

	idx = 0
	_, err = jsonparser.ArrayEach(c,
		func(value []byte, dt jsonparser.ValueType, offset int, err error) {
			if dt != jsonparser.Object {
				errCb = errors.New("value is not object")
				return
			}

			idxStr := fmt.Sprintf("[%d]", idx)
			var indexDecVal string
			indexDecVal, errCb = jsonparser.GetString(value, "INDEX_DEC")
			if errCb != nil {
				return
			}

			var renameTo string
			switch indexDecVal {
			case "历史变更-企业名称变更":
				renameTo = "企业名称变更"
			case "历史变更-地址变更":
				renameTo = "地址变更"
			case "变更的风险提示\n（减资）":
				renameTo = "是否减资"
			}

			if renameTo != "" {
				toSet := fmt.Sprintf(`"%s"`, renameTo)
				c1, errCb := jsonparser.Set(*s.content, []byte(toSet), append(paths, idxStr, "INDEX_DEC")...)
				if errCb != nil {
					return
				}
				*s.content = c1
			}
			idx++
		}, paths...)

	if errCb != nil {
		return errCb
	}
	if err != nil {
		return err
	}

	r, err := s.summary()
	if err != nil {
		return errors.Wrap(err, "error from ClaRiskIndexes.summary")
	}
	summaryBytes, err := json.Marshal(r)
	if err != nil {
		return err
	}
	tempC, err := jsonparser.Set(*s.content, summaryBytes, "impExpEntReport", "riskIndexesSummary")
	if err != nil {
		return err
	}

	*s.content = tempC

	return nil
}

func (s *ClaRiskIndexes) lambdaReformatIndexValue(template string, value string) string {
	if value == "0" || value == "-" {
		return "无"
	} else if _, err := strconv.ParseFloat(value, 32); err != nil {
		return value
	} else {
		return fmt.Sprintf(template, value)
	}
}

func (s *ClaRiskIndexes) templateValueReformat() map[string]string {
	return map[string]string{
		"历史变更-企业名称变更":   "变更企业名称%s次",
		"历史变更-地址变更":     "变更地址%s次",
		"变更的风险提示\n（减资）": "减资注册资本%s次",
	}
}

func (s *ClaRiskIndexes) round(value string) string {
	if v, err := strconv.ParseFloat(value, 64); err != nil {
		return fmt.Sprintf("%d", int(math.Round(v)))
	} else {
		return value
	}
}

type riskIndexSummarySta struct {
	Total     *int `json:"total"`
	Normal    *int `json:"normal"`
	Attention *int `json:"attention"`
	Abnormal  *int `json:"abnormal"`
}

type riskIndexSummary struct {
	Total       *int      `json:"total"`
	Normal      *int      `json:"normal"`
	Attention   *int      `json:"attention"`
	Abnormal    *int      `json:"abnormal"`
	AbnormalDec *[]string `json:"abnormalTags"`
}

func (s *ClaRiskIndexes) summary() (*riskIndexSummary, error) {
	model := models.RcRiskIndex{}
	db := sdk.Runtime.GetDbByKey(model.TableName())
	sta := riskIndexSummarySta{}
	err := db.Table(model.TableName()).
		Raw(
			`select count(*)                                             as total,
					   sum(case when index_flag = '正常' then 1 else 0 end) as normal,
					   sum(case when index_flag = '关注' then 1 else 0 end) as attention,
					   sum(case when index_flag = '异常' then 1 else 0 end) as abnormal
				from rc_risk_index
				where content_id = ?`, s.contentId).
		First(&sta).
		Error
	if err != nil {
		return nil, err
	}
	abnormalDec := make([]string, 0)

	err = db.Table(model.TableName()).
		Select("distinct index_dec").
		Where("index_flag = ?", "异常").
		Where("content_id = ?", s.contentId).
		Pluck("index_dec", &abnormalDec).
		Error
	if err != nil {
		return nil, err
	}

	r := riskIndexSummary{
		Total:       sta.Total,
		Normal:      sta.Normal,
		Attention:   sta.Attention,
		Abnormal:    sta.Abnormal,
		AbnormalDec: &abnormalDec,
	}

	return &r, nil
}

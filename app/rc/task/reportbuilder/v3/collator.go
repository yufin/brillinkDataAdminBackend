package v3

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

type ClaRiskIndexes struct {
}

func (s *ClaRiskIndexes) Collating(content *[]byte, contentId int64) error {
	
	return nil
}

func (s *ClaRiskIndexes) CollatingWasted(content *[]byte) error {
	paths := []string{"impExpEntReport", "riskIndexes"}
	var idx int
	var errCb, err error
	c := *content

	{
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
					fmt.Println(indexValue, "fixme")
				case "历史变更-地址变更":
					resetVal = "地址变更"
				case "变更的风险提示\n（减资）":
					resetVal = "是否减资"
				}

				if resetVal != "" {
					toSet := fmt.Sprintf(`"%s"`, resetVal)
					*content, errCb = jsonparser.Set(*content, []byte(toSet), append(paths, idxStr, "INDEX_DEC")...)
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
	}

	{
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
					*content, errCb = jsonparser.Set(*content, []byte(toSet), append(paths, idxStr, "INDEX_DEC")...)
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
	}

	return nil
}

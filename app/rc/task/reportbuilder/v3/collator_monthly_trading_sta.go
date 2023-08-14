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
	"strings"
)

type ClaMonthlyTradingSta struct {
	content   *[]byte
	contentId int64
}
type staResult struct {
	ContentId int64   `json:"contentId"`
	Year      string  `json:"year"`
	Result    float64 `json:"result"`
}

type summaries struct {
	S1 []string `json:"summary1"`
	S2 []string `json:"summary2"`
	S3 []string `json:"summary3"`
}

type ydxsInfo struct {
	Cyl      string `json:"CYL"`
	Cyl1     string `json:"CYL_1"`
	Cyl2     string `json:"CYL_2"`
	FpXxAvg  string `json:"FP_XX_AVG"`
	FpXxAvg1 string `json:"FP_XX_AVG_1"`
	FpXxAvg2 string `json:"FP_XX_AVG_2"`
	SbAvg    string `json:"SB_AVG"`
	SbAvg1   string `json:"SB_AVG_1"`
	SbAvg2   string `json:"SB_AVG_2"`
	Title    string `json:"TITLE"`
	Title1   string `json:"TITLE_1"`
	Title2   string `json:"TITLE_2"`
}

type ydcgInfo struct {
	GncgAvg1 string `json:"GNCG_AVG1"`
	GncgAvg2 string `json:"GNCG_AVG2"`
	GncgAvg3 string `json:"GNCG_AVG3"`
	HjAvg1   string `json:"HJ_AVG1"`
	HjAvg2   string `json:"HJ_AVG2"`
	HjAvg3   string `json:"HJ_AVG3"`
	JkcgAvg1 string `json:"JKCG_AVG1"`
	JkcgAvg2 string `json:"JKCG_AVG2"`
	JkcgAvg3 string `json:"JKCG_AVG3"`
	Title1   string `json:"TITLE1"`
	Title2   string `json:"TITLE2"`
	Title3   string `json:"TITLE3"`
}

type valueRange struct {
	Min float64
	Max float64
}

func (s *ClaMonthlyTradingSta) SetContent(content *[]byte, contentId int64) {
	s.content = content
	s.contentId = contentId
}

func (s *ClaMonthlyTradingSta) Collating() error {
	//paths := []string{"impExpEntReport", "ydxsqkInfo"}
	var err error

	model := models.RcMonthlySales{}
	db := sdk.Runtime.GetDbByKey(model.TableName())
	res := make([]staResult, 0)
	err = db.Table(model.TableName()).
		Raw(`SELECT x.content_id, x.year, 
       				CASE 
       				    WHEN x.sbkjzsr = 0 THEN (x.fpkjsr - y.hj_m) / x.fpkjsr * 100 
       					ELSE  (x.sbkjzsr - y.hj_m) / x.sbkjzsr * 100 
       			  	END  AS result
					FROM (
						SELECT content_id, SUBSTRING(attributed_month, 1, 4) AS year, SUM(sbkjzsr) AS sbkjzsr, SUM(fpkjsr) AS fpkjsr
						FROM rc_monthly_sales
							WHERE content_id=?
						GROUP BY content_id, SUBSTRING(attributed_month, 1, 4)
					) x
					JOIN (
						SELECT content_id, SUBSTRING(attributed_month, 1, 4) AS year, SUM(hj_m) AS hj_m
						FROM rc_monthly_purchase
							WHERE content_id=?
						GROUP BY content_id, SUBSTRING(attributed_month, 1, 4)
					) y ON x.content_id = y.content_id AND x.year = y.year
					ORDER BY year desc ;`, s.contentId, s.contentId).
		Scan(&res).
		Error
	if err != nil {
		return err
	}
	resBytes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	tempC, err := jsonparser.Set(*s.content, resBytes, "impExpEntReport", "ydxsqkInfo", "czxsPercentage")
	if err != nil {
		return err
	}
	*s.content = tempC

	var xsSummary, cgSummary, lastTwoYearSummary, lastYearSummary, thisYearSummary string

	xsSummary, err = s.subSummary1()
	if err != nil {
		return err
	}
	cgSummary, err = s.subSummary2()
	if err != nil {
		return err
	}

	summaries1 := []string{
		xsSummary,
		cgSummary,
	}

	summary := fmt.Sprintf("%s年-%s年差值/销售金额分别是%.2f%%，%.2f%%，%.2f%%；", res[2].Year, res[0].Year, res[2].Result, res[1].Result, res[0].Result)
	summaries2 := []string{
		summary,
	}

	var summaries3 []string
	var sepcialRule1, sepcialRule2, sepcialRule3, sepcialRule4 string
	if s.subRule(res[2].Result, res[0].Result) {
		sepcialRule1, err = s.subSummary4(res, 0, 1)
		if err != nil {
			return err
		}
		summaries3 = []string{
			sepcialRule1,
		}
	} else if s.subRule(res[2].Result, res[1].Result) {
		sepcialRule2, err = s.subSummary4(res, 1, 2)
		if err != nil {
			return err
		}
		sepcialRule4, err = s.subSummary4(res, 0, 4)
		if err != nil {
			return err
		}
		summaries3 = []string{
			sepcialRule2,
			sepcialRule4,
		}
	} else if s.subRule(res[1].Result, res[0].Result) {
		sepcialRule3, err = s.subSummary4(res, 0, 3)
		if err != nil {
			return err
		}
		sepcialRule4, err = s.subSummary4(res, 2, 4)
		if err != nil {
			return err
		}
		summaries3 = []string{
			sepcialRule4,
			sepcialRule3,
		}
	} else {
		lastTwoYearSummary, err = s.subSummary4(res, 2, 4)
		lastYearSummary, err = s.subSummary4(res, 1, 4)
		thisYearSummary, err = s.subSummary4(res, 0, 4)

		summaries3 = []string{
			lastTwoYearSummary,
			lastYearSummary,
			thisYearSummary,
		}
	}

	summries := summaries{
		S1: summaries1,
		S2: summaries2,
		S3: summaries3,
	}
	summriesBytes, err := json.Marshal(summries)
	if err != nil {
		return err
	}

	c, err := jsonparser.Set(*s.content, summriesBytes, "impExpEntReport", "ydxscgSummary")
	if err != nil {
		return err
	}
	*s.content = c

	if err != nil {
		return err
	}
	return nil
}

func (s *ClaMonthlyTradingSta) subSummary1() (string, error) {
	c := *s.content
	infoBytes, dt, _, err := jsonparser.Get(c, "impExpEntReport", "ydxsqkInfo")
	if err != nil {
		return "", err
	}
	if dt != jsonparser.Object {
		return "", errors.New("ydxsqkInfo not object")
	}
	info := ydxsInfo{}
	err = json.Unmarshal(infoBytes, &info)
	if err != nil {
		return "", err
	}
	var lastTwoYearXsVal, lastYearXsVal, thisYearXsVal float64

	lastTwoYearXsTime := info.Title2[:4]
	lastYearXsTime := info.Title1[:4]
	thisYearXsTime := info.Title[:4]

	formaterSbAvg2 := strings.ReplaceAll(info.SbAvg2, ",", "")
	Sbavg2, err := strconv.ParseFloat(formaterSbAvg2, 64)
	if err != nil {
		Sbavg2 = 0
	}

	formaterSbAvg1 := strings.ReplaceAll(info.SbAvg1, ",", "")
	Sbavg1, err := strconv.ParseFloat(formaterSbAvg1, 64)
	if err != nil {
		Sbavg1 = 0
	}

	formaterSbAvg := strings.ReplaceAll(info.SbAvg, ",", "")
	Sbavg, err := strconv.ParseFloat(formaterSbAvg, 64)
	if err != nil {
		Sbavg = 0
	}

	if Sbavg2 != 0 {
		lastTwoYearXsVal = Sbavg2
	} else {
		formaterFpXxAvg2 := strings.ReplaceAll(info.FpXxAvg2, ",", "")
		lastTwoYearXsVal, err = strconv.ParseFloat(formaterFpXxAvg2, 64)
		if err != nil {
			lastTwoYearXsVal = 0
		}
	}
	if Sbavg1 != 0 {
		lastYearXsVal = Sbavg1
	} else {
		formaterFpXxAvg1 := strings.ReplaceAll(info.FpXxAvg1, ",", "")
		lastYearXsVal, err = strconv.ParseFloat(formaterFpXxAvg1, 64)
		if err != nil {
			lastYearXsVal = 0
		}
	}
	if Sbavg != 0 {
		thisYearXsVal = Sbavg
	} else {
		formaterFpXxAvg := strings.ReplaceAll(info.FpXxAvg, ",", "")
		thisYearXsVal, err = strconv.ParseFloat(formaterFpXxAvg, 64)
		if err != nil {
			thisYearXsVal = 0
		}
	}

	var xsAvg1, xsAvg2, xsAvg3 float64
	if lastTwoYearXsVal != 0 {
		xsAvg1 = (lastYearXsVal - lastTwoYearXsVal) / lastTwoYearXsVal * 100
	}
	xsAvgDesc1, err := s.subSummary3(xsAvg1)
	if err != nil {
		return "", err
	}

	if lastYearXsVal != 0 {
		xsAvg2 = (thisYearXsVal - lastYearXsVal) / lastYearXsVal * 100
	}
	xsAvgDesc2, err := s.subSummary3(xsAvg2)
	if err != nil {
		return "", err
	}

	if lastTwoYearXsVal != 0 {
		xsAvg3 = (thisYearXsVal - lastTwoYearXsVal) / lastTwoYearXsVal * 100
	}
	xsAvgDesc3, err := s.subSummary3(xsAvg3)
	if err != nil {
		return "", err
	}

	lastTwoYearXsVal = math.Round(lastTwoYearXsVal)
	lastYearXsVal = math.Round(lastYearXsVal)
	thisYearXsVal = math.Round(thisYearXsVal)

	Summary := fmt.Sprintf("%s年-%s年每月平均销售额分别是%.f、%.f、%.f；其中%s年-%s年%s，%s年-%s年%s，整体来看，%s年-%s年%s。", lastTwoYearXsTime, thisYearXsTime, lastTwoYearXsVal, lastYearXsVal, thisYearXsVal, lastTwoYearXsTime, lastYearXsTime, xsAvgDesc1, lastYearXsTime, thisYearXsTime, xsAvgDesc2, lastTwoYearXsTime, thisYearXsTime, xsAvgDesc3)

	return Summary, nil

}

func (s *ClaMonthlyTradingSta) subSummary2() (string, error) {
	c := *s.content
	infoBytes, dt, _, err := jsonparser.Get(c, "impExpEntReport", "ydcgqkInfo")
	if err != nil {
		return "", err
	}
	if dt != jsonparser.Object {
		return "", errors.New("ydcgqkInfo not object")
	}
	info := ydcgInfo{}
	err = json.Unmarshal(infoBytes, &info)
	if err != nil {
		return "", err
	}
	var lastTwoYearCgVal, lastYearCgVal, thisYearCgVal float64

	lastTwoYearCgTime := info.Title1[:4]
	lastYearCgTime := info.Title2[:4]
	thisYearCgTime := info.Title3[:4]

	formaterHjAvg1 := strings.ReplaceAll(info.HjAvg1, ",", "")
	lastTwoYearCgVal, err = strconv.ParseFloat(formaterHjAvg1, 64)
	if err != nil {
		lastTwoYearCgVal = 0
	}

	formaterHjAvg2 := strings.ReplaceAll(info.HjAvg2, ",", "")
	lastYearCgVal, err = strconv.ParseFloat(formaterHjAvg2, 64)
	if err != nil {
		lastYearCgVal = 0
	}

	formaterHjAvg3 := strings.ReplaceAll(info.HjAvg3, ",", "")
	thisYearCgVal, err = strconv.ParseFloat(formaterHjAvg3, 64)
	if err != nil {
		thisYearCgVal = 0
	}

	var cgAvg1, cgAvg2, cgAvg3 float64
	if lastTwoYearCgVal != 0 {
		cgAvg1 = (lastYearCgVal - lastTwoYearCgVal) / lastTwoYearCgVal * 100
	}
	cgAvgDesc1, err := s.subSummary3(cgAvg1)
	if err != nil {
		return "", err
	}

	if lastYearCgVal != 0 {
		cgAvg2 = (thisYearCgVal - lastYearCgVal) / lastYearCgVal * 100
	}
	cgAvgDesc2, err := s.subSummary3(cgAvg2)
	if err != nil {
		return "", err
	}

	if lastTwoYearCgVal != 0 {
		cgAvg3 = (thisYearCgVal - lastTwoYearCgVal) / lastTwoYearCgVal * 100
	}
	cgAvgDesc3, err := s.subSummary3(cgAvg3)
	if err != nil {
		return "", err
	}

	lastTwoYearCgVal = math.Round(lastTwoYearCgVal)
	lastYearCgVal = math.Round(lastYearCgVal)
	thisYearCgVal = math.Round(thisYearCgVal)

	Summary := fmt.Sprintf("%s年-%s年每月平均销售额分别是%.f、%.f、%.f；其中%s年-%s年%s，%s年-%s年%s，整体来看，%s年-%s年%s。", lastTwoYearCgTime, thisYearCgTime, lastTwoYearCgVal, lastYearCgVal, thisYearCgVal, lastTwoYearCgTime, lastYearCgTime, cgAvgDesc1, lastYearCgTime, thisYearCgTime, cgAvgDesc2, lastTwoYearCgTime, thisYearCgTime, cgAvgDesc3)

	return Summary, nil

}

func (s *ClaMonthlyTradingSta) subSummary3(value float64) (string, error) {
	var Summary string

	if value > 30 {
		Summary = fmt.Sprintf("<font color= #EF5644>增加%.2f%%</font>，<font color= #EF5644>快速增长</font>", value)
	} else if value > 10 && value <= 30 {
		Summary = fmt.Sprintf("<font color= #EF5644>增加%.2f%%</font>，<font color= #EF5644>小幅增长</font>", value)
	} else if value >= 0 && value <= 10 {
		Summary = fmt.Sprintf("<font color= #EF5644>增加%.2f%%</font>，<font color= #E6A23C>趋势平稳</font>", value)
	} else if value >= -10 && value < 0 {
		Summary = fmt.Sprintf("<font color= #67C23A>减少%.2f%%</font>，<font color= #E6A23C>趋势平稳</font>", value)
	} else if value >= -30 && value < -10 {
		Summary = fmt.Sprintf("<font color= #67C23A>减少%.2f%%</font>，<font color= #67C23A>小幅下降</font>", value)
	} else if value < -30 {
		Summary = fmt.Sprintf("<font color= #67C23A>减少%.2f%%</font>，<font color= #67C23A>大幅下降</font>", value)
	}

	return Summary, nil

}

func (s *ClaMonthlyTradingSta) subSummary4(res []staResult, index int, specialRule int) (string, error) {
	var Summary string

	resetVal := res[index].Result
	switch specialRule {
	case 1:
		if resetVal > 30 {
			Summary = fmt.Sprintf("%s年-%s年利润较高，备货很少。", res[2].Year, res[0].Year)
		} else if resetVal > 10 && resetVal <= 30 {
			Summary = fmt.Sprintf("%s年-%s年有部分利润，备货较少。", res[2].Year, res[0].Year)
		} else if resetVal >= -10 && resetVal <= 10 {
			Summary = fmt.Sprintf("%s年-%s年基本上按销定采，供应链管理能力较好。", res[2].Year, res[0].Year)
		} else if resetVal >= -30 && resetVal < -10 {
			Summary = fmt.Sprintf("%s年-%s年小幅度备货或者部分原材料呆滞，需关注库存风险。", res[2].Year, res[0].Year)
		} else if resetVal < -30 {
			Summary = fmt.Sprintf("%s年-%s年大幅度备货或大部分原材料呆滞，需重点关注库存风险。", res[2].Year, res[0].Year)
		}
	case 2:
		if resetVal > 30 {
			Summary = fmt.Sprintf("%s年-%s年利润较高，备货很少。", res[2].Year, res[1].Year)
		} else if resetVal > 10 && resetVal <= 30 {
			Summary = fmt.Sprintf("%s年-%s年有部分利润，备货较少。", res[2].Year, res[1].Year)
		} else if resetVal >= -10 && resetVal <= 10 {
			Summary = fmt.Sprintf("%s年-%s年基本上按销定采，供应链管理能力较好。", res[2].Year, res[1].Year)
		} else if resetVal >= -30 && resetVal < -10 {
			Summary = fmt.Sprintf("%s年-%s年小幅度备货或者部分原材料呆滞，需关注库存风险。", res[2].Year, res[1].Year)
		} else if resetVal < -30 {
			Summary = fmt.Sprintf("%s年-%s年大幅度备货或大部分原材料呆滞，需重点关注库存风险。", res[2].Year, res[1].Year)
		}
	case 3:
		if resetVal > 30 {
			Summary = fmt.Sprintf("%s年-%s年利润较高，备货很少。", res[1].Year, res[0].Year)
		} else if resetVal > 10 && resetVal <= 30 {
			Summary = fmt.Sprintf("%s年-%s年有部分利润，备货较少。", res[1].Year, res[0].Year)
		} else if resetVal >= -10 && resetVal <= 10 {
			Summary = fmt.Sprintf("%s年-%s年基本上按销定采，供应链管理能力较好。", res[1].Year, res[0].Year)
		} else if resetVal >= -30 && resetVal < -10 {
			Summary = fmt.Sprintf("%s年-%s年小幅度备货或者部分原材料呆滞，需关注库存风险。", res[1].Year, res[0].Year)
		} else if resetVal < -30 {
			Summary = fmt.Sprintf("%s年-%s年大幅度备货或大部分原材料呆滞，需重点关注库存风险。", res[1].Year, res[0].Year)
		}
	case 4:
		if resetVal > 30 {
			Summary = fmt.Sprintf("%s年利润较高，备货很少。", res[index].Year)
		} else if resetVal > 10 && resetVal <= 30 {
			Summary = fmt.Sprintf("%s年有部分利润，备货较少。", res[index].Year)
		} else if resetVal >= -10 && resetVal <= 10 {
			Summary = fmt.Sprintf("%s年基本上按销定采，供应链管理能力较好。", res[index].Year)
		} else if resetVal >= -30 && resetVal < -10 {
			Summary = fmt.Sprintf("%s年小幅度备货或者部分原材料呆滞，需关注库存风险。", res[index].Year)
		} else if resetVal < -30 {
			Summary = fmt.Sprintf("%s年大幅度备货或大部分原材料呆滞，需重点关注库存风险。", res[index].Year)
		}
	}
	return Summary, nil
}

func (s *ClaMonthlyTradingSta) subRule(value1 float64, value2 float64) bool {
	valueRanges := []valueRange{
		{Min: math.Inf(-1), Max: -30},
		{Min: -30, Max: -10},
		{Min: -10, Max: 0},
		{Min: 0, Max: 10},
		{Min: 10, Max: 30},
		{Min: 30, Max: math.Inf(1)},
	}

	for _, r := range valueRanges {
		if value1 >= r.Min && value1 < r.Max && value2 >= r.Min && value2 < r.Max {
			return true // 区间范围相同
		}
	}

	return false
}

package v3

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/go-admin-team/go-admin-core/sdk"
	"github.com/pkg/errors"
	"go-admin/app/rc/models"
)

// ClaFinancialSummary
type ClaFinancialSummary struct {
	content   *[]byte
	contentId int64
}

func (s *ClaFinancialSummary) SetContent(content *[]byte, contentId int64) {
	s.content = content
	s.contentId = contentId
}

func (s *ClaFinancialSummary) Collating() error {
	summary, err := s.revenueSummary()
	if err != nil {
		return errors.Wrap(err, "revenueSummary error")
	}

	sb, err := json.Marshal(summary)
	if err != nil {
		return errors.Wrap(err, "json Marshal error")
	}
	c, err := jsonparser.Set(*s.content, sb, "impExpEntReport", "lrbAnalysisSummary")
	*s.content = c
	return nil
}

func (s *ClaFinancialSummary) revenueSummary() ([]string, error) {
	res := make([]string, 0)

	tbRrd := models.RcRevenueDetail{}
	db := sdk.Runtime.GetDbByKey(tbRrd.TableName())

	income := make([]models.RcRevenueDetail, 0)
	err := db.Model(models.RcRevenueDetail{}).
		Where("field = ?", "营业收入").
		Where("content_id = ?", s.contentId).
		Where("val is not null").
		Order("period_start asc").
		Scan(&income).
		Error
	if err != nil {
		return []string{}, err
	}
	if len(income) == 0 {
		return []string{}, nil
	}

	t0 := income[0].PeriodStart
	t1 := income[1].PeriodStart
	t2 := income[2].PeriodEnd
	val0, _ := income[0].Val.Decimal.Float64()
	val1, _ := income[1].Val.Decimal.Float64()
	val2, _ := income[2].Val.Decimal.Float64()

	sOverView := fmt.Sprintf(`%d年-%s营业收入是%s, %s, %s;`,
		t0.Year(),
		t2.Format("2006年01月"),
		income[2].Val.Decimal.String(),
		income[1].Val.Decimal.String(),
		income[0].Val.Decimal.String())

	res = append(res, sOverView)

	desc1 := "-"
	desc2 := "-"
	rangeDesc1 := "-"
	rangeDesc2 := "-"
	s1 := "-"
	s2 := "-"
	//s0 := "-"

	var rate1, rate2 float64

	avgIncome2 := "-"
	if income[2].Val.Valid {
		avgIncomeF := val2 / float64(t2.YearDay()) * 30
		avgIncome2 = fmt.Sprintf("%.2f万元", avgIncomeF)
	}

	if income[0].Val.Valid && income[1].Val.Valid {
		desc1, rate1 = s.evalRate(val0, val1)
		rangeDesc1 = s.pctRangeDesc(rate1)
	}
	if income[1].Val.Valid && income[2].Val.Valid {
		desc2, rate2 = s.evalRate(val1, val2)
		rangeDesc2 = s.pctRangeDesc(rate2)
	}

	s2 = fmt.Sprintf(
		`其中%d年年初至%s（最近）月均收入%s，同比%s，%s。`,
		t2.Year(),
		t2.Format("01月02日"),
		avgIncome2,
		desc2,
		//t2.Year(),
		rangeDesc2)
	res = append(res, s2)

	s1 = fmt.Sprintf(
		`%d年全年收入同比%s，%s`,
		t1.Year(),
		desc1,
		//t1.Year(),
		rangeDesc1,
	)
	res = append(res, s1)

	//s0 = fmt.Sprintf("%d年%s，净利润为、   、   ；毛利率为    、   、    ；净利率为    、   、    ；近期公司XX;",
	//	t0.Year(),
	//)
	//res = append(res, s0)
	return res, err
}

func (s *ClaFinancialSummary) evalRate(lastVal float64, recentVal float64) (string, float64) {
	if lastVal == 0 {
		return "", 0
	}
	pct := (recentVal - lastVal) / lastVal * 100
	if pct > 0 {
		return fmt.Sprintf("增长%.2f%%", pct), pct
	} else if pct < 0 {
		return fmt.Sprintf("下降%.2f%%", pct), pct
	} else {
		return "持平", pct
	}
}

func (s *ClaFinancialSummary) pctRangeDesc(pct float64) string {
	//正数：(10%,30%]营业收入小幅度增长，(30%,100%]营业收入大幅度增长，(100%,170%]
	//营业收入约增长一倍，(170%,270%]营业收入约增长2倍，以此类比：（A70%-B70%）营业收入约增长B倍；
	//负数(10%,40%]营业收入小幅度下滑，(40%,60%]营业收入约下滑一半，(60%，100%]营业收入大幅度下滑

	if pct > 10 && pct <= 30 {
		return "小幅度增长"
	} else if pct > 30 && pct <= 100 {
		return "大幅度增长"
	} else if pct > 100 && pct <= 170 {
		return "约增长一倍"
	} else if pct > 170 {
		x := int((pct-70)/100) + 1
		return fmt.Sprintf("约增长%d倍", x)
	} else if pct < -10 && pct >= -40 {
		return "小幅度下降"
	} else if pct < -40 && pct >= -60 {
		return "约下降一半"
	} else if pct < -60 && pct > -100 {
		return "大幅度下滑"
	} else {
		return "几乎持平"
	}

}

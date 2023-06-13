package models

import (
	"github.com/shopspring/decimal"
	"go-admin/common/models"
)

// RcDecisionParam
type RcDecisionParam struct {
	models.Model
	ContentId               int64               `json:"-" gorm:"comment:参数id" xlsx:"content_id"`
	SwSdsnbCyrs             *int                `json:"sw_sdsnb_cyrs" gorm:"column:sw_sdsnb_cyrs; comment:从业人数" xlsx:"sw_sdsnb_cyrs"`
	GsGdct                  *int                `json:"gs_gdct" gorm:"column:gs_gdct; comment:股东穿透" xlsx:"gs_gdct"`
	GsGdwdx                 *int                `json:"gs_gdwdx" gorm:"column:gs_gdwdx; comment:股东稳定性" xlsx:"gs_gdwdx"`
	GsFrwdx                 *int                `json:"gs_frwdx" gorm:"column:gs_frwdx; comment:法人稳定性" xlsx:"gs_frwdx"`
	LhCylwz                 *int                `json:"lh_cylwz" gorm:"column:lh_cylwz; comment:产业链价值度" xlsx:"lh_cylwz"`
	LhMdPpjzl               *int                `json:"lh_md_ppjzl" gorm:"column:lh_md_ppjzl; comment:技术/品牌竞争力" xlsx:"lh_md_ppjzl"`
	MdQybq                  *int                `json:"md_qybq" gorm:"column:md_qybq; comment:企业标签" xlsx:"md_qybq"`
	SwCwbbYyzjzzts          *int                `json:"sw_cwbb_yyzjzzts" gorm:"column:sw_cwbb_yyzjzzts; comment:营运资金周转天数" xlsx:"sw_cwbb_yyzjzzts"`
	SfFhSsqkQy              *int                `json:"sf_fh_ssqk_qy" gorm:"column:sf_fh_ssqk_qy; comment:历史诉讼情况" xlsx:"sf_fh_ssqk_qy"`
	SwJcxxNsrxypj           *string             `json:"sw_jcxx_nsrxypj" gorm:"column:sw_jcxx_nsrxypj; comment:纳税信用评级" xlsx:"sw_jcxx_nsrxypj"`
	ZxYhsxqk                *int                `json:"zx_yhsxqk" gorm:"column:zx_yhsxqk; comment:银行授信情况" xlsx:"zx_yhsxqk"`
	ZxDsfsxqk               *int                `json:"zx_dsfsxqk" gorm:"column:zx_dsfsxqk; comment:第三方机构授信情况" xlsx:"zx_dsfsxqk"`
	LhQylx                  *int                `json:"lh_qylx" gorm:"column:lh_qylx; comment:企业类型" xlsx:"lh_qylx"`
	Nsrsbh                  *string             `json:"nsrsbh" gorm:"column:nsrsbh; comment:纳税人识别号" xlsx:"nsrsbh"`
	SwSbNszeZzsqysds12m     decimal.NullDecimal `json:"sw_sb_nsze_zzsqysds_12m" gorm:"type:decimal(18,10);column:sw_sb_nsze_zzsqysds_12m; comment:近12个月企业缴税规模(增值税+所得税)" xlsx:"sw_sb_nsze_zzsqysds_12m"`
	SwSbNszezzlZzsqysds12mA decimal.NullDecimal `json:"sw_sb_nszezzl_zzsqysds_12m_a" gorm:"type:decimal(18,10);column:sw_sb_nszezzl_zzsqysds_12m_a; comment:近12个月纳税总额增长率(增值税+企业所得税)(环比、同比)" xlsx:"sw_sb_nszezzl_zzsqysds_12m_a"`
	SwSdsnbGzxjzzjezzl      decimal.NullDecimal `json:"sw_sdsnb_gzxjzzjezzl" gorm:"type:decimal(18,10);column:sw_sdsnb_gzxjzzjezzl; comment:工资薪金支出账载金额增长率" xlsx:"sw_sdsnb_gzxjzzjezzl"`
	SwSbzsSflhypld12m       decimal.NullDecimal `json:"sw_sbzs_sflhypld_12m" gorm:"type:decimal(18,10);column:sw_sbzs_sflhypld_12m; comment:近12个月税负率行业偏离度" xlsx:"sw_sbzs_sflhypld_12m"`
	SwSdsnbYjfy             decimal.NullDecimal `json:"sw_sdsnb_yjfy" gorm:"type:decimal(18,10);column:sw_sdsnb_yjfy; comment:研究费用" xlsx:"sw_sdsnb_yjfy"`
	FpJxLxfy12m             decimal.NullDecimal `json:"fp_jx_lxfy_12m" gorm:"type:decimal(18,10);column:fp_jx_lxfy_12m; comment:近12个月利息费用" xlsx:"fp_jx_lxfy_12m"`
	SwCwbbSszb              decimal.NullDecimal `json:"sw_cwbb_sszb" gorm:"type:decimal(18,10);column:sw_cwbb_sszb; comment:(当年)实收资本(或股本)" xlsx:"sw_cwbb_sszb"`
	FpJySychjeZb12mLh       decimal.NullDecimal `json:"fp_jy_sychje_zb_12m_lh" gorm:"type:decimal(18,10);column:fp_jy_sychje_zb_12m_lh; comment:供应商稳定性(重要供应商金额占比重合 程度)近12个月与前12个月" xlsx:"fp_jy_sychje_zb_12m_lh"`
	FpJxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_jx_zyjyjezb_12m_lh" gorm:"type:decimal(18,10);column:fp_jx_zyjyjezb_12m_lh; comment:优质供应商占比(含500强、上市公司 等)按照撞库销售额" xlsx:"fp_jx_zyjyjezb_12m_lh"`
	FpXxXychjeZb12mLh       decimal.NullDecimal `json:"fp_xx_xychje_zb_12m_lh" gorm:"type:decimal(18,10);column:fp_xx_xychje_zb_12m_lh; comment:客户稳定性(重要客户金额占比重合程 度)近12个月与前12个月" xlsx:"fp_xx_xychje_zb_12m_lh"`
	FpXxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_xx_zyjyjezb_12m_lh" gorm:"type:decimal(18,10);column:fp_xx_zyjyjezb_12m_lh; comment:优质客户占比(含500强、上市公司等) 按照撞库销售额" xlsx:"fp_xx_zyjyjezb_12m_lh"`
	SwSbQbxse12m            decimal.NullDecimal `json:"sw_sb_qbxse_12m" gorm:"type:decimal(18,10);column:sw_sb_qbxse_12m; comment:年销售额" xlsx:"sw_sb_qbxse_12m"`
	SwSbQbxsezzl12m         decimal.NullDecimal `json:"sw_sb_qbxsezzl_12m" gorm:"type:decimal(18,10);column:sw_sb_qbxsezzl_12m; comment:年销售增长率" xlsx:"sw_sb_qbxsezzl_12m"`
	SwSbLsxs12m             decimal.NullDecimal `json:"sw_sb_lsxs_12m" gorm:"type:decimal(18,10);column:sw_sb_lsxs_12m; comment:近12个月全部销售额变异系数" xlsx:"sw_sb_lsxs_12m"`
	SwCwbbChzztsCb          decimal.NullDecimal `json:"sw_cwbb_chzzts_cb" gorm:"type:decimal(18,10);column:sw_cwbb_chzzts_cb; comment:库存风险" xlsx:"sw_cwbb_chzzts_cb"`
	SwCwbbZcfzl             decimal.NullDecimal `json:"sw_cwbb_zcfzl" gorm:"type:decimal(18,10);column:sw_cwbb_zcfzl; comment:资产负债率" xlsx:"sw_cwbb_zcfzl"`
	SwCwbbMlrzzlv           decimal.NullDecimal `json:"sw_cwbb_mlrzzlv" gorm:"type:decimal(18,10);column:sw_cwbb_mlrzzlv; comment:毛利润增长率" xlsx:"sw_cwbb_mlrzzlv"`
	SwCwbbJlrzzlv           decimal.NullDecimal `json:"sw_cwbb_jlrzzlv" gorm:"type:decimal(18,10);column:sw_cwbb_jlrzzlv; comment:净利润增长率" xlsx:"sw_cwbb_jlrzzlv"`
	SwCwbbJzcszlv           decimal.NullDecimal `json:"sw_cwbb_jzcszlv" gorm:"type:decimal(18,10);column:sw_cwbb_jzcszlv; comment:净资产收益率" xlsx:"sw_cwbb_jzcszlv"`
	SwJcxxClnx              decimal.NullDecimal `json:"sw_jcxx_clnx" gorm:"type:decimal(18,10);column:sw_jcxx_clnx; comment:经营年限(年份数)" xlsx:"sw_jcxx_clnx"`
	StatusCode              int                 `json:"statusCode" gorm:"comment:状态码;" xlsx:"StatusCode"`
	models.ModelTime
	models.ControlBy
}

func (*RcDecisionParam) TableName() string {
	return "rc_decision_param"
}

func (e *RcDecisionParam) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RcDecisionParam) GetId() interface{} {
	return e.Id
}

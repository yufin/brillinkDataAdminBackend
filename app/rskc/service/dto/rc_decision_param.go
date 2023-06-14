package dto

import (
	"github.com/shopspring/decimal"
	"go-admin/app/rskc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
	"go-admin/utils"
)

type RcDecisionParamGetPageReq struct {
	dto.Pagination          `search:"-"`
	Id                      int64               `form:"id"  search:"type:exact;column:id;table:rc_decision_param" comment:"id"`
	ContentId               int64               `form:"contentId"  search:"type:exact;column:content_id;table:rc_decision_param" comment:""`
	SwSdsnbCyrs             *int                `form:"sw_sdsnb_cyrs"  search:"type:exact;column:sw_sdsnb_cyrs;table:rc_decision_param" comment:""`
	GsGdct                  *int                `form:"gs_gdct"  search:"type:exact;column:gs_gdct;table:rc_decision_param" comment:""`
	GsGdwdx                 *int                `form:"gs_gdwdx"  search:"type:exact;column:gs_gdwdx;table:rc_decision_param" comment:""`
	GsFrwdx                 *int                `form:"gs_frwdx"  search:"type:exact;column:gs_frwdx;table:rc_decision_param" comment:""`
	LhCylwz                 *int                `form:"lh_cylwz"  search:"type:exact;column:lh_cylwz;table:rc_decision_param" comment:""`
	LhMdPpjzl               *int                `form:"lh_md_ppjzl"  search:"type:exact;column:lh_md_ppjzl;table:rc_decision_param" comment:""`
	MdQybq                  *int                `form:"md_qybq"  search:"type:exact;column:md_qybq;table:rc_decision_param" comment:""`
	SwCwbbYyzjzzts          *int                `form:"sw_cwbb_yyzjzzts"  search:"type:exact;column:sw_cwbb_yyzjzzts;table:rc_decision_param" comment:""`
	SfFhSsqkQy              *int                `form:"sf_fh_ssqk_qy"  search:"type:exact;column:sf_fh_ssqk_qy;table:rc_decision_param" comment:"纳税信用评级"`
	SwJcxxNsrxypj           *string             `form:"sw_jcxx_nsrxypj"  search:"type:exact;column:sw_jcxx_nsrxypj;table:rc_decision_param" comment:""`
	ZxYhsxqk                *int                `form:"zx_yhsxqk"  search:"type:exact;column:zx_yhsxqk;table:rc_decision_param" comment:""`
	ZxDsfsxqk               *int                `form:"zx_dsfsxqk"  search:"type:exact;column:zx_dsfsxqk;table:rc_decision_param" comment:""`
	LhQylx                  *int                `form:"lh_qylx"  search:"type:exact;column:lh_qylx;table:rc_decision_param" comment:""`
	Nsrsbh                  *string             `form:"nsrsbh"  search:"type:exact;column:nsrsbh;table:rc_decision_param" comment:""`
	SwSbNszeZzsqysds12m     decimal.NullDecimal `form:"sw_sb_nsze_zzsqysds_12m"  search:"type:exact;column:sw_sb_nsze_zzsqysds_12m;table:rc_decision_param" comment:""`
	SwSbNszezzlZzsqysds12mA decimal.NullDecimal `form:"sw_sb_nszezzl_zzsqysds_12m_a"  search:"type:exact;column:sw_sb_nszezzl_zzsqysds_12m_a;table:rc_decision_param" comment:""`
	SwSdsnbGzxjzzjezzl      decimal.NullDecimal `form:"sw_sdsnb_gzxjzzjezzl"  search:"type:exact;column:sw_sdsnb_gzxjzzjezzl;table:rc_decision_param" comment:""`
	SwSbzsSflhypld12m       decimal.NullDecimal `form:"sw_sbzs_sflhypld_12m"  search:"type:exact;column:sw_sbzs_sflhypld_12m;table:rc_decision_param" comment:""`
	SwSdsnbYjfy             decimal.NullDecimal `form:"sw_sdsnb_yjfy"  search:"type:exact;column:sw_sdsnb_yjfy;table:rc_decision_param" comment:""`
	FpJxLxfy12m             decimal.NullDecimal `form:"fp_jx_lxfy_12m"  search:"type:exact;column:fp_jx_lxfy_12m;table:rc_decision_param" comment:""`
	SwCwbbSszb              decimal.NullDecimal `form:"sw_cwbb_sszb"  search:"type:exact;column:sw_cwbb_sszb;table:rc_decision_param" comment:""`
	FpJySychjeZb12mLh       decimal.NullDecimal `form:"fp_jy_sychje_zb_12m_lh"  search:"type:exact;column:fp_jy_sychje_zb_12m_lh;table:rc_decision_param" comment:""`
	FpJxZyjyjezb12mLh       decimal.NullDecimal `form:"fp_jx_zyjyjezb_12m_lh"  search:"type:exact;column:fp_jx_zyjyjezb_12m_lh;table:rc_decision_param" comment:""`
	FpXxXychjeZb12mLh       decimal.NullDecimal `form:"fp_xx_xychje_zb_12m_lh"  search:"type:exact;column:fp_xx_xychje_zb_12m_lh;table:rc_decision_param" comment:""`
	FpXxZyjyjezb12mLh       decimal.NullDecimal `form:"fp_xx_zyjyjezb_12m_lh"  search:"type:exact;column:fp_xx_zyjyjezb_12m_lh;table:rc_decision_param" comment:""`
	SwSbQbxse12m            decimal.NullDecimal `form:"sw_sb_qbxse_12m"  search:"type:exact;column:sw_sb_qbxse_12m;table:rc_decision_param" comment:""`
	SwSbQbxsezzl12m         decimal.NullDecimal `form:"sw_sb_qbxsezzl_12m"  search:"type:exact;column:sw_sb_qbxsezzl_12m;table:rc_decision_param" comment:""`
	SwSbLsxs12m             decimal.NullDecimal `form:"sw_sb_lsxs_12m"  search:"type:exact;column:sw_sb_lsxs_12m;table:rc_decision_param" comment:""`
	SwCwbbChzztsCb          decimal.NullDecimal `form:"sw_cwbb_chzzts_cb"  search:"type:exact;column:sw_cwbb_chzzts_cb;table:rc_decision_param" comment:""`
	SwCwbbZcfzl             decimal.NullDecimal `form:"sw_cwbb_zcfzl"  search:"type:exact;column:sw_cwbb_zcfzl;table:rc_decision_param" comment:""`
	SwCwbbMlrzzlv           decimal.NullDecimal `form:"sw_cwbb_mlrzzlv"  search:"type:exact;column:sw_cwbb_mlrzzlv;table:rc_decision_param" comment:""`
	SwCwbbJlrzzlv           decimal.NullDecimal `form:"sw_cwbb_jlrzzlv"  search:"type:exact;column:sw_cwbb_jlrzzlv;table:rc_decision_param" comment:""`
	SwCwbbJzcszlv           decimal.NullDecimal `form:"sw_cwbb_jzcszlv"  search:"type:exact;column:sw_cwbb_jzcszlv;table:rc_decision_param" comment:""`
	SwJcxxClnx              decimal.NullDecimal `form:"sw_jcxx_clnx"  search:"type:exact;column:sw_jcxx_clnx;table:rc_decision_param" comment:""`
	StatusCode              int                 `form:"statusCode"  search:"type:exact;column:status_code;table:rc_decision_param" comment:""`
	StartTime               string              `form:"startTime" search:"type:gte;column:created_at;table:rc_decision_param" comment:"创建时间"`
	EndTime                 string              `form:"endTime" search:"type:lte;column:created_at;table:rc_decision_param" comment:"创建时间"`
	RcDecisionParamPageOrder
}

type RcDecisionParamPageOrder struct {
	Id                      string `form:"idOrder"  search:"type:order;column:id;table:rc_decision_param"`
	SwSdsnbCyrs             string `form:"swSdsnbCyrsOrder"  search:"type:order;column:sw_sdsnb_cyrs;table:rc_decision_param"`
	GsGdct                  string `form:"gsGdctOrder"  search:"type:order;column:gs_gdct;table:rc_decision_param"`
	GsGdwdx                 string `form:"gsGdwdxOrder"  search:"type:order;column:gs_gdwdx;table:rc_decision_param"`
	GsFrwdx                 string `form:"gsFrwdxOrder"  search:"type:order;column:gs_frwdx;table:rc_decision_param"`
	LhCylwz                 string `form:"lhCylwzOrder"  search:"type:order;column:lh_cylwz;table:rc_decision_param"`
	LhMdPpjzl               string `form:"lhMdPpjzlOrder"  search:"type:order;column:lh_md_ppjzl;table:rc_decision_param"`
	MdQybq                  string `form:"mdQybqOrder"  search:"type:order;column:md_qybq;table:rc_decision_param"`
	SwCwbbYyzjzzts          string `form:"swCwbbYyzjzztsOrder"  search:"type:order;column:sw_cwbb_yyzjzzts;table:rc_decision_param"`
	SfFhSsqkQy              string `form:"sfFhSsqkQyOrder"  search:"type:order;column:sf_fh_ssqk_qy;table:rc_decision_param"`
	SwJcxxNsrxypj           string `form:"swJcxxNsrxypjOrder"  search:"type:order;column:sw_jcxx_nsrxypj;table:rc_decision_param"`
	ZxYhsxqk                string `form:"zxYhsxqkOrder"  search:"type:order;column:zx_yhsxqk;table:rc_decision_param"`
	ZxDsfsxqk               string `form:"zxDsfsxqkOrder"  search:"type:order;column:zx_dsfsxqk;table:rc_decision_param"`
	LhQylx                  string `form:"lhQylxOrder"  search:"type:order;column:lh_qylx;table:rc_decision_param"`
	Nsrsbh                  string `form:"nsrsbhOrder"  search:"type:order;column:nsrsbh;table:rc_decision_param"`
	SwSbNszeZzsqysds12m     string `form:"swSbNszeZzsqysds12mOrder"  search:"type:order;column:sw_sb_nsze_zzsqysds_12m;table:rc_decision_param"`
	SwSbNszezzlZzsqysds12mA string `form:"swSbNszezzlZzsqysds12mAOrder"  search:"type:order;column:sw_sb_nszezzl_zzsqysds_12m_a;table:rc_decision_param"`
	SwSdsnbGzxjzzjezzl      string `form:"swSdsnbGzxjzzjezzlOrder"  search:"type:order;column:sw_sdsnb_gzxjzzjezzl;table:rc_decision_param"`
	SwSbzsSflhypld12m       string `form:"swSbzsSflhypld12mOrder"  search:"type:order;column:sw_sbzs_sflhypld_12m;table:rc_decision_param"`
	SwSdsnbYjfy             string `form:"swSdsnbYjfyOrder"  search:"type:order;column:sw_sdsnb_yjfy;table:rc_decision_param"`
	FpJxLxfy12m             string `form:"fpJxLxfy12mOrder"  search:"type:order;column:fp_jx_lxfy_12m;table:rc_decision_param"`
	SwCwbbSszb              string `form:"swCwbbSszbOrder"  search:"type:order;column:sw_cwbb_sszb;table:rc_decision_param"`
	FpJySychjeZb12mLh       string `form:"fpJySychjeZb12mLhOrder"  search:"type:order;column:fp_jy_sychje_zb_12m_lh;table:rc_decision_param"`
	FpJxZyjyjezb12mLh       string `form:"fpJxZyjyjezb12mLhOrder"  search:"type:order;column:fp_jx_zyjyjezb_12m_lh;table:rc_decision_param"`
	FpXxXychjeZb12mLh       string `form:"fpXxXychjeZb12mLhOrder"  search:"type:order;column:fp_xx_xychje_zb_12m_lh;table:rc_decision_param"`
	FpXxZyjyjezb12mLh       string `form:"fpXxZyjyjezb12mLhOrder"  search:"type:order;column:fp_xx_zyjyjezb_12m_lh;table:rc_decision_param"`
	SwSbQbxse12m            string `form:"swSbQbxse12mOrder"  search:"type:order;column:sw_sb_qbxse_12m;table:rc_decision_param"`
	SwSbQbxsezzl12m         string `form:"swSbQbxsezzl12mOrder"  search:"type:order;column:sw_sb_qbxsezzl_12m;table:rc_decision_param"`
	SwSbLsxs12m             string `form:"swSbLsxs12mOrder"  search:"type:order;column:sw_sb_lsxs_12m;table:rc_decision_param"`
	SwCwbbChzztsCb          string `form:"swCwbbChzztsCbOrder"  search:"type:order;column:sw_cwbb_chzzts_cb;table:rc_decision_param"`
	SwCwbbZcfzl             string `form:"swCwbbZcfzlOrder"  search:"type:order;column:sw_cwbb_zcfzl;table:rc_decision_param"`
	SwCwbbMlrzzlv           string `form:"swCwbbMlrzzlvOrder"  search:"type:order;column:sw_cwbb_mlrzzlv;table:rc_decision_param"`
	SwCwbbJlrzzlv           string `form:"swCwbbJlrzzlvOrder"  search:"type:order;column:sw_cwbb_jlrzzlv;table:rc_decision_param"`
	SwCwbbJzcszlv           string `form:"swCwbbJzcszlvOrder"  search:"type:order;column:sw_cwbb_jzcszlv;table:rc_decision_param"`
	SwJcxxClnx              string `form:"swJcxxClnxOrder"  search:"type:order;column:sw_jcxx_clnx;table:rc_decision_param"`
	StatusCode              string `form:"statusCodeOrder"  search:"type:order;column:status_code;table:rc_decision_param"`
}

func (m *RcDecisionParamGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RcDecisionParamGetResp struct {
	Id                      int64               `json:"id"`                           // id
	SwSdsnbCyrs             *int                `json:"sw_sdsnb_cyrs"`                //
	GsGdct                  *int                `json:"gs_gdct"`                      //
	GsGdwdx                 *int                `json:"gs_gdwdx"`                     //
	GsFrwdx                 *int                `json:"gs_frwdx"`                     //
	LhCylwz                 *int                `json:"lh_cylwz"`                     //
	LhMdPpjzl               *int                `json:"lh_md_ppjzl"`                  //
	MdQybq                  *int                `json:"md_qybq"`                      //
	SwCwbbYyzjzzts          *int                `json:"sw_cwbb_yyzjzzts"`             //
	SfFhSsqkQy              *int                `json:"sf_fh_ssqk_qy"`                // 纳税信用评级
	SwJcxxNsrxypj           *string             `json:"sw_jcxx_nsrxypj"`              //
	ZxYhsxqk                *int                `json:"zx_yhsxqk"`                    //
	ZxDsfsxqk               *int                `json:"zx_dsfsxqk"`                   //
	LhQylx                  *int                `json:"lh_qylx"`                      //
	Nsrsbh                  *string             `json:"nsrsbh"`                       //
	SwSbNszeZzsqysds12m     decimal.NullDecimal `json:"sw_sb_nsze_zzsqysds_12m"`      //
	SwSbNszezzlZzsqysds12mA decimal.NullDecimal `json:"sw_sb_nszezzl_zzsqysds_12m_a"` //
	SwSdsnbGzxjzzjezzl      decimal.NullDecimal `json:"sw_sdsnb_gzxjzzjezzl"`         //
	SwSbzsSflhypld12m       decimal.NullDecimal `json:"sw_sbzs_sflhypld_12m"`         //
	SwSdsnbYjfy             decimal.NullDecimal `json:"sw_sdsnb_yjfy"`                //
	FpJxLxfy12m             decimal.NullDecimal `json:"fp_jx_lxfy_12m"`               //
	SwCwbbSszb              decimal.NullDecimal `json:"sw_cwbb_sszb"`                 //
	FpJySychjeZb12mLh       decimal.NullDecimal `json:"fp_jy_sychje_zb_12m_lh"`       //
	FpJxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_jx_zyjyjezb_12m_lh"`        //
	FpXxXychjeZb12mLh       decimal.NullDecimal `json:"fp_xx_xychje_zb_12m_lh"`       //
	FpXxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_xx_zyjyjezb_12m_lh"`        //
	SwSbQbxse12m            decimal.NullDecimal `json:"sw_sb_qbxse_12m"`              //
	SwSbQbxsezzl12m         decimal.NullDecimal `json:"sw_sb_qbxsezzl_12m"`           //
	SwSbLsxs12m             decimal.NullDecimal `json:"sw_sb_lsxs_12m"`               //
	SwCwbbChzztsCb          decimal.NullDecimal `json:"sw_cwbb_chzzts_cb"`            //
	SwCwbbZcfzl             decimal.NullDecimal `json:"sw_cwbb_zcfzl"`                //
	SwCwbbMlrzzlv           decimal.NullDecimal `json:"sw_cwbb_mlrzzlv"`              //
	SwCwbbJlrzzlv           decimal.NullDecimal `json:"sw_cwbb_jlrzzlv"`              //
	SwCwbbJzcszlv           decimal.NullDecimal `json:"sw_cwbb_jzcszlv"`              //
	SwJcxxClnx              decimal.NullDecimal `json:"sw_jcxx_clnx"`                 //
	StatusCode              int                 `json:"statusCode"`                   //
	common.ControlBy
}

func (s *RcDecisionParamGetResp) Generate(model *models.RcDecisionParam) {
	if s.Id == 0 {
		s.Id = model.Id
	}
	s.SwSdsnbCyrs = model.SwSdsnbCyrs
	s.GsGdct = model.GsGdct
	s.GsGdwdx = model.GsGdwdx
	s.GsFrwdx = model.GsFrwdx
	s.LhCylwz = model.LhCylwz
	s.LhMdPpjzl = model.LhMdPpjzl
	s.MdQybq = model.MdQybq
	s.SwCwbbYyzjzzts = model.SwCwbbYyzjzzts
	s.SfFhSsqkQy = model.SfFhSsqkQy
	s.SwJcxxNsrxypj = model.SwJcxxNsrxypj
	s.ZxYhsxqk = model.ZxYhsxqk
	s.ZxDsfsxqk = model.ZxDsfsxqk
	s.LhQylx = model.LhQylx
	s.Nsrsbh = model.Nsrsbh
	s.SwSbNszeZzsqysds12m = model.SwSbNszeZzsqysds12m
	s.SwSbNszezzlZzsqysds12mA = model.SwSbNszezzlZzsqysds12mA
	s.SwSdsnbGzxjzzjezzl = model.SwSdsnbGzxjzzjezzl
	s.SwSbzsSflhypld12m = model.SwSbzsSflhypld12m
	s.SwSdsnbYjfy = model.SwSdsnbYjfy
	s.FpJxLxfy12m = model.FpJxLxfy12m
	s.SwCwbbSszb = model.SwCwbbSszb
	s.FpJySychjeZb12mLh = model.FpJySychjeZb12mLh
	s.FpJxZyjyjezb12mLh = model.FpJxZyjyjezb12mLh
	s.FpXxXychjeZb12mLh = model.FpXxXychjeZb12mLh
	s.FpXxZyjyjezb12mLh = model.FpXxZyjyjezb12mLh
	s.SwSbQbxse12m = model.SwSbQbxse12m
	s.SwSbQbxsezzl12m = model.SwSbQbxsezzl12m
	s.SwSbLsxs12m = model.SwSbLsxs12m
	s.SwCwbbChzztsCb = model.SwCwbbChzztsCb
	s.SwCwbbZcfzl = model.SwCwbbZcfzl
	s.SwCwbbMlrzzlv = model.SwCwbbMlrzzlv
	s.SwCwbbJlrzzlv = model.SwCwbbJlrzzlv
	s.SwCwbbJzcszlv = model.SwCwbbJzcszlv
	s.SwJcxxClnx = model.SwJcxxClnx
	s.StatusCode = model.StatusCode
	s.CreateBy = model.CreateBy
}

type RcDecisionParamInsertReq struct {
	Id                      int64               `json:"-"`                            // id
	ContentId               int64               `json:"contentId"`                    //
	SwSdsnbCyrs             *int                `json:"sw_sdsnb_cyrs"`                //
	GsGdct                  *int                `json:"gs_gdct"`                      //
	GsGdwdx                 *int                `json:"gs_gdwdx"`                     //
	GsFrwdx                 *int                `json:"gs_frwdx"`                     //
	LhCylwz                 *int                `json:"lh_cylwz"`                     //
	LhMdPpjzl               *int                `json:"lh_md_ppjzl"`                  //
	MdQybq                  *int                `json:"md_qybq"`                      //
	SwCwbbYyzjzzts          *int                `json:"sw_cwbb_yyzjzzts"`             //
	SfFhSsqkQy              *int                `json:"sf_fh_ssqk_qy"`                // 纳税信用评级
	SwJcxxNsrxypj           *string             `json:"sw_jcxx_nsrxypj"`              //
	ZxYhsxqk                *int                `json:"zx_yhsxqk"`                    //
	ZxDsfsxqk               *int                `json:"zx_dsfsxqk"`                   //
	LhQylx                  *int                `json:"lh_qylx"`                      //
	Nsrsbh                  *string             `json:"nsrsbh"`                       //
	SwSbNszeZzsqysds12m     decimal.NullDecimal `json:"sw_sb_nsze_zzsqysds_12m"`      //
	SwSbNszezzlZzsqysds12mA decimal.NullDecimal `json:"sw_sb_nszezzl_zzsqysds_12m_a"` //
	SwSdsnbGzxjzzjezzl      decimal.NullDecimal `json:"sw_sdsnb_gzxjzzjezzl"`         //
	SwSbzsSflhypld12m       decimal.NullDecimal `json:"sw_sbzs_sflhypld_12m"`         //
	SwSdsnbYjfy             decimal.NullDecimal `json:"sw_sdsnb_yjfy"`                //
	FpJxLxfy12m             decimal.NullDecimal `json:"fp_jx_lxfy_12m"`               //
	SwCwbbSszb              decimal.NullDecimal `json:"sw_cwbb_sszb"`                 //
	FpJySychjeZb12mLh       decimal.NullDecimal `json:"fp_jy_sychje_zb_12m_lh"`       //
	FpJxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_jx_zyjyjezb_12m_lh"`        //
	FpXxXychjeZb12mLh       decimal.NullDecimal `json:"fp_xx_xychje_zb_12m_lh"`       //
	FpXxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_xx_zyjyjezb_12m_lh"`        //
	SwSbQbxse12m            decimal.NullDecimal `json:"sw_sb_qbxse_12m"`              //
	SwSbQbxsezzl12m         decimal.NullDecimal `json:"sw_sb_qbxsezzl_12m"`           //
	SwSbLsxs12m             decimal.NullDecimal `json:"sw_sb_lsxs_12m"`               //
	SwCwbbChzztsCb          decimal.NullDecimal `json:"sw_cwbb_chzzts_cb"`            //
	SwCwbbZcfzl             decimal.NullDecimal `json:"sw_cwbb_zcfzl"`                //
	SwCwbbMlrzzlv           decimal.NullDecimal `json:"sw_cwbb_mlrzzlv"`              //
	SwCwbbJlrzzlv           decimal.NullDecimal `json:"sw_cwbb_jlrzzlv"`              //
	SwCwbbJzcszlv           decimal.NullDecimal `json:"sw_cwbb_jzcszlv"`              //
	SwJcxxClnx              decimal.NullDecimal `json:"sw_jcxx_clnx"`                 //
	StatusCode              int                 `json:"statusCode"`                   //
	common.ControlBy
}

func (s *RcDecisionParamInsertReq) Generate(model *models.RcDecisionParam) {
	if s.Id == 0 {
		model.Model = common.Model{Id: utils.NewFlakeId()}
	}
	model.ContentId = s.ContentId
	model.SwSdsnbCyrs = s.SwSdsnbCyrs
	model.GsGdct = s.GsGdct
	model.GsGdwdx = s.GsGdwdx
	model.GsFrwdx = s.GsFrwdx
	model.LhCylwz = s.LhCylwz
	model.LhMdPpjzl = s.LhMdPpjzl
	model.MdQybq = s.MdQybq
	model.SwCwbbYyzjzzts = s.SwCwbbYyzjzzts
	model.SfFhSsqkQy = s.SfFhSsqkQy
	model.SwJcxxNsrxypj = s.SwJcxxNsrxypj
	model.ZxYhsxqk = s.ZxYhsxqk
	model.ZxDsfsxqk = s.ZxDsfsxqk
	model.LhQylx = s.LhQylx
	model.Nsrsbh = s.Nsrsbh
	model.SwSbNszeZzsqysds12m = s.SwSbNszeZzsqysds12m
	model.SwSbNszezzlZzsqysds12mA = s.SwSbNszezzlZzsqysds12mA
	model.SwSdsnbGzxjzzjezzl = s.SwSdsnbGzxjzzjezzl
	model.SwSbzsSflhypld12m = s.SwSbzsSflhypld12m
	model.SwSdsnbYjfy = s.SwSdsnbYjfy
	model.FpJxLxfy12m = s.FpJxLxfy12m
	model.SwCwbbSszb = s.SwCwbbSszb
	model.FpJySychjeZb12mLh = s.FpJySychjeZb12mLh
	model.FpJxZyjyjezb12mLh = s.FpJxZyjyjezb12mLh
	model.FpXxXychjeZb12mLh = s.FpXxXychjeZb12mLh
	model.FpXxZyjyjezb12mLh = s.FpXxZyjyjezb12mLh
	model.SwSbQbxse12m = s.SwSbQbxse12m
	model.SwSbQbxsezzl12m = s.SwSbQbxsezzl12m
	model.SwSbLsxs12m = s.SwSbLsxs12m
	model.SwCwbbChzztsCb = s.SwCwbbChzztsCb
	model.SwCwbbZcfzl = s.SwCwbbZcfzl
	model.SwCwbbMlrzzlv = s.SwCwbbMlrzzlv
	model.SwCwbbJlrzzlv = s.SwCwbbJlrzzlv
	model.SwCwbbJzcszlv = s.SwCwbbJzcszlv
	model.SwJcxxClnx = s.SwJcxxClnx
	model.StatusCode = s.StatusCode
	model.CreateBy = s.CreateBy
}

func (s *RcDecisionParamInsertReq) GetId() interface{} {
	return s.Id
}

type RcDecisionParamUpdateReq struct {
	Id                      int64               `uri:"id"`                            // id
	ContentId               int64               `json:"contentId"`                    //
	SwSdsnbCyrs             *int                `json:"sw_sdsnb_cyrs"`                //
	GsGdct                  *int                `json:"gs_gdct"`                      //
	GsGdwdx                 *int                `json:"gs_gdwdx"`                     //
	GsFrwdx                 *int                `json:"gs_frwdx"`                     //
	LhCylwz                 *int                `json:"lh_cylwz"`                     //
	LhMdPpjzl               *int                `json:"lh_md_ppjzl"`                  //
	MdQybq                  *int                `json:"md_qybq"`                      //
	SwCwbbYyzjzzts          *int                `json:"sw_cwbb_yyzjzzts"`             //
	SfFhSsqkQy              *int                `json:"sf_fh_ssqk_qy"`                // 纳税信用评级
	SwJcxxNsrxypj           *string             `json:"sw_jcxx_nsrxypj"`              //
	ZxYhsxqk                *int                `json:"zx_yhsxqk"`                    //
	ZxDsfsxqk               *int                `json:"zx_dsfsxqk"`                   //
	LhQylx                  *int                `json:"lh_qylx"`                      //
	Nsrsbh                  *string             `json:"nsrsbh"`                       //
	SwSbNszeZzsqysds12m     decimal.NullDecimal `json:"sw_sb_nsze_zzsqysds_12m"`      //
	SwSbNszezzlZzsqysds12mA decimal.NullDecimal `json:"sw_sb_nszezzl_zzsqysds_12m_a"` //
	SwSdsnbGzxjzzjezzl      decimal.NullDecimal `json:"sw_sdsnb_gzxjzzjezzl"`         //
	SwSbzsSflhypld12m       decimal.NullDecimal `json:"sw_sbzs_sflhypld_12m"`         //
	SwSdsnbYjfy             decimal.NullDecimal `json:"sw_sdsnb_yjfy"`                //
	FpJxLxfy12m             decimal.NullDecimal `json:"fp_jx_lxfy_12m"`               //
	SwCwbbSszb              decimal.NullDecimal `json:"sw_cwbb_sszb"`                 //
	FpJySychjeZb12mLh       decimal.NullDecimal `json:"fp_jy_sychje_zb_12m_lh"`       //
	FpJxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_jx_zyjyjezb_12m_lh"`        //
	FpXxXychjeZb12mLh       decimal.NullDecimal `json:"fp_xx_xychje_zb_12m_lh"`       //
	FpXxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_xx_zyjyjezb_12m_lh"`        //
	SwSbQbxse12m            decimal.NullDecimal `json:"sw_sb_qbxse_12m"`              //
	SwSbQbxsezzl12m         decimal.NullDecimal `json:"sw_sb_qbxsezzl_12m"`           //
	SwSbLsxs12m             decimal.NullDecimal `json:"sw_sb_lsxs_12m"`               //
	SwCwbbChzztsCb          decimal.NullDecimal `json:"sw_cwbb_chzzts_cb"`            //
	SwCwbbZcfzl             decimal.NullDecimal `json:"sw_cwbb_zcfzl"`                //
	SwCwbbMlrzzlv           decimal.NullDecimal `json:"sw_cwbb_mlrzzlv"`              //
	SwCwbbJlrzzlv           decimal.NullDecimal `json:"sw_cwbb_jlrzzlv"`              //
	SwCwbbJzcszlv           decimal.NullDecimal `json:"sw_cwbb_jzcszlv"`              //
	SwJcxxClnx              decimal.NullDecimal `json:"sw_jcxx_clnx"`                 //
	StatusCode              int                 `json:"statusCode"`                   //
	common.ControlBy
}

func (s *RcDecisionParamUpdateReq) Generate(model *models.RcDecisionParam) {
	if s.Id == 0 {
		model.Model = common.Model{Id: utils.NewFlakeId()}
	}
	model.SwSdsnbCyrs = s.SwSdsnbCyrs
	model.GsGdct = s.GsGdct
	model.GsGdwdx = s.GsGdwdx
	model.GsFrwdx = s.GsFrwdx
	model.LhCylwz = s.LhCylwz
	model.LhMdPpjzl = s.LhMdPpjzl
	model.MdQybq = s.MdQybq
	model.SwCwbbYyzjzzts = s.SwCwbbYyzjzzts
	model.SfFhSsqkQy = s.SfFhSsqkQy
	model.SwJcxxNsrxypj = s.SwJcxxNsrxypj
	model.ZxYhsxqk = s.ZxYhsxqk
	model.ZxDsfsxqk = s.ZxDsfsxqk
	model.LhQylx = s.LhQylx
	model.Nsrsbh = s.Nsrsbh
	model.SwSbNszeZzsqysds12m = s.SwSbNszeZzsqysds12m
	model.SwSbNszezzlZzsqysds12mA = s.SwSbNszezzlZzsqysds12mA
	model.SwSdsnbGzxjzzjezzl = s.SwSdsnbGzxjzzjezzl
	model.SwSbzsSflhypld12m = s.SwSbzsSflhypld12m
	model.SwSdsnbYjfy = s.SwSdsnbYjfy
	model.FpJxLxfy12m = s.FpJxLxfy12m
	model.SwCwbbSszb = s.SwCwbbSszb
	model.FpJySychjeZb12mLh = s.FpJySychjeZb12mLh
	model.FpJxZyjyjezb12mLh = s.FpJxZyjyjezb12mLh
	model.FpXxXychjeZb12mLh = s.FpXxXychjeZb12mLh
	model.FpXxZyjyjezb12mLh = s.FpXxZyjyjezb12mLh
	model.SwSbQbxse12m = s.SwSbQbxse12m
	model.SwSbQbxsezzl12m = s.SwSbQbxsezzl12m
	model.SwSbLsxs12m = s.SwSbLsxs12m
	model.SwCwbbChzztsCb = s.SwCwbbChzztsCb
	model.SwCwbbZcfzl = s.SwCwbbZcfzl
	model.SwCwbbMlrzzlv = s.SwCwbbMlrzzlv
	model.SwCwbbJlrzzlv = s.SwCwbbJlrzzlv
	model.SwCwbbJzcszlv = s.SwCwbbJzcszlv
	model.SwJcxxClnx = s.SwJcxxClnx
	model.StatusCode = s.StatusCode
	model.UpdateBy = s.UpdateBy
}

func (s *RcDecisionParamUpdateReq) GetId() interface{} {
	return s.Id
}

// RcDecisionParamGetReq 功能获取请求参数
type RcDecisionParamGetReq struct {
	Id int64 `uri:"id"`
}

func (s *RcDecisionParamGetReq) GetId() interface{} {
	return s.Id
}

// RcDecisionParamDeleteReq 功能删除请求参数
type RcDecisionParamDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RcDecisionParamDeleteReq) GetId() interface{} {
	return s.Ids
}

type RcDecisionParamExport struct {
	Id                      int64               `json:"id" gorm:"primaryKey;comment:id" xlsx:"id"`
	ContentId               int64               `json:"contentId" gorm:"comment:contentId" xlsx:"contentId"`
	SwSdsnbCyrs             *int                `json:"sw_sdsnb_cyrs" gorm:"comment:SwSdsnbCyrs" xlsx:"SwSdsnbCyrs"`
	GsGdct                  *int                `json:"gs_gdct" gorm:"comment:GsGdct" xlsx:"GsGdct"`
	GsGdwdx                 *int                `json:"gs_gdwdx" gorm:"comment:GsGdwdx" xlsx:"GsGdwdx"`
	GsFrwdx                 *int                `json:"gs_frwdx" gorm:"comment:GsFrwdx" xlsx:"GsFrwdx"`
	LhCylwz                 *int                `json:"lh_cylwz" gorm:"comment:LhCylwz" xlsx:"LhCylwz"`
	LhMdPpjzl               *int                `json:"lh_md_ppjzl" gorm:"comment:LhMdPpjzl" xlsx:"LhMdPpjzl"`
	MdQybq                  *int                `json:"md_qybq" gorm:"comment:MdQybq" xlsx:"MdQybq"`
	SwCwbbYyzjzzts          *int                `json:"sw_cwbb_yyzjzzts" gorm:"comment:SwCwbbYyzjzzts" xlsx:"SwCwbbYyzjzzts"`
	SfFhSsqkQy              *int                `json:"sf_fh_ssqk_qy" gorm:"comment:纳税信用评级" xlsx:"纳税信用评级"`
	SwJcxxNsrxypj           *string             `json:"sw_jcxx_nsrxypj" gorm:"comment:SwJcxxNsrxypj" xlsx:"SwJcxxNsrxypj"`
	ZxYhsxqk                *int                `json:"zx_yhsxqk" gorm:"comment:ZxYhsxqk" xlsx:"ZxYhsxqk"`
	ZxDsfsxqk               *int                `json:"zx_dsfsxqk" gorm:"comment:ZxDsfsxqk" xlsx:"ZxDsfsxqk"`
	LhQylx                  *int                `json:"lh_qylx" gorm:"comment:LhQylx" xlsx:"LhQylx"`
	Nsrsbh                  *string             `json:"nsrsbh" gorm:"comment:Nsrsbh" xlsx:"Nsrsbh"`
	SwSbNszeZzsqysds12m     decimal.NullDecimal `json:"sw_sb_nsze_zzsqysds_12m" gorm:"comment:SwSbNszeZzsqysds12m" xlsx:"SwSbNszeZzsqysds12m"`
	SwSbNszezzlZzsqysds12mA decimal.NullDecimal `json:"sw_sb_nszezzl_zzsqysds_12m_a" gorm:"comment:SwSbNszezzlZzsqysds12mA" xlsx:"SwSbNszezzlZzsqysds12mA"`
	SwSdsnbGzxjzzjezzl      decimal.NullDecimal `json:"sw_sdsnb_gzxjzzjezzl" gorm:"comment:SwSdsnbGzxjzzjezzl" xlsx:"SwSdsnbGzxjzzjezzl"`
	SwSbzsSflhypld12m       decimal.NullDecimal `json:"sw_sbzs_sflhypld_12m" gorm:"comment:SwSbzsSflhypld12m" xlsx:"SwSbzsSflhypld12m"`
	SwSdsnbYjfy             decimal.NullDecimal `json:"sw_sdsnb_yjfy" gorm:"comment:SwSdsnbYjfy" xlsx:"SwSdsnbYjfy"`
	FpJxLxfy12m             decimal.NullDecimal `json:"fp_jx_lxfy_12m" gorm:"comment:FpJxLxfy12m" xlsx:"FpJxLxfy12m"`
	SwCwbbSszb              decimal.NullDecimal `json:"sw_cwbb_sszb" gorm:"comment:SwCwbbSszb" xlsx:"SwCwbbSszb"`
	FpJySychjeZb12mLh       decimal.NullDecimal `json:"fp_jy_sychje_zb_12m_lh" gorm:"comment:FpJySychjeZb12mLh" xlsx:"FpJySychjeZb12mLh"`
	FpJxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_jx_zyjyjezb_12m_lh" gorm:"comment:FpJxZyjyjezb12mLh" xlsx:"FpJxZyjyjezb12mLh"`
	FpXxXychjeZb12mLh       decimal.NullDecimal `json:"fp_xx_xychje_zb_12m_lh" gorm:"comment:FpXxXychjeZb12mLh" xlsx:"FpXxXychjeZb12mLh"`
	FpXxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_xx_zyjyjezb_12m_lh" gorm:"comment:FpXxZyjyjezb12mLh" xlsx:"FpXxZyjyjezb12mLh"`
	SwSbQbxse12m            decimal.NullDecimal `json:"sw_sb_qbxse_12m" gorm:"comment:SwSbQbxse12m" xlsx:"SwSbQbxse12m"`
	SwSbQbxsezzl12m         decimal.NullDecimal `json:"sw_sb_qbxsezzl_12m" gorm:"comment:SwSbQbxsezzl12m" xlsx:"SwSbQbxsezzl12m"`
	SwSbLsxs12m             decimal.NullDecimal `json:"sw_sb_lsxs_12m" gorm:"comment:SwSbLsxs12m" xlsx:"SwSbLsxs12m"`
	SwCwbbChzztsCb          decimal.NullDecimal `json:"sw_cwbb_chzzts_cb" gorm:"comment:SwCwbbChzztsCb" xlsx:"SwCwbbChzztsCb"`
	SwCwbbZcfzl             decimal.NullDecimal `json:"sw_cwbb_zcfzl" gorm:"comment:SwCwbbZcfzl" xlsx:"SwCwbbZcfzl"`
	SwCwbbMlrzzlv           decimal.NullDecimal `json:"sw_cwbb_mlrzzlv" gorm:"comment:SwCwbbMlrzzlv" xlsx:"SwCwbbMlrzzlv"`
	SwCwbbJlrzzlv           decimal.NullDecimal `json:"sw_cwbb_jlrzzlv" gorm:"comment:SwCwbbJlrzzlv" xlsx:"SwCwbbJlrzzlv"`
	SwCwbbJzcszlv           decimal.NullDecimal `json:"sw_cwbb_jzcszlv" gorm:"comment:SwCwbbJzcszlv" xlsx:"SwCwbbJzcszlv"`
	SwJcxxClnx              decimal.NullDecimal `json:"sw_jcxx_clnx" gorm:"comment:SwJcxxClnx" xlsx:"SwJcxxClnx"`
	StatusCode              int                 `json:"statusCode" gorm:"comment:StatusCode" xlsx:"StatusCode"`
}

type RcDecisionParamDecisionRequestBody struct {
	Id                      int64               `uri:"-"`                             // id
	ContentId               int64               `json:"-"`                            //
	SwSdsnbCyrs             *int                `json:"sw_sdsnb_cyrs"`                //
	GsGdct                  *int                `json:"gs_gdct"`                      //
	GsGdwdx                 *int                `json:"gs_gdwdx"`                     //
	GsFrwdx                 *int                `json:"gs_frwdx"`                     //
	LhCylwz                 *int                `json:"lh_cylwz"`                     //
	LhMdPpjzl               *int                `json:"lh_md_ppjzl"`                  //
	MdQybq                  *int                `json:"md_qybq"`                      //
	SwCwbbYyzjzzts          *int                `json:"sw_cwbb_yyzjzzts"`             //
	SfFhSsqkQy              *int                `json:"sf_fh_ssqk_qy"`                // 纳税信用评级
	SwJcxxNsrxypj           *string             `json:"sw_jcxx_nsrxypj"`              //
	ZxYhsxqk                *int                `json:"zx_yhsxqk"`                    //
	ZxDsfsxqk               *int                `json:"zx_dsfsxqk"`                   //
	LhQylx                  *int                `json:"lh_qylx"`                      //
	Nsrsbh                  *string             `json:"nsrsbh"`                       //
	SwSbNszeZzsqysds12m     decimal.NullDecimal `json:"sw_sb_nsze_zzsqysds_12m"`      //
	SwSbNszezzlZzsqysds12mA decimal.NullDecimal `json:"sw_sb_nszezzl_zzsqysds_12m_a"` //
	SwSdsnbGzxjzzjezzl      decimal.NullDecimal `json:"sw_sdsnb_gzxjzzjezzl"`         //
	SwSbzsSflhypld12m       decimal.NullDecimal `json:"sw_sbzs_sflhypld_12m"`         //
	SwSdsnbYjfy             decimal.NullDecimal `json:"sw_sdsnb_yjfy"`                //
	FpJxLxfy12m             decimal.NullDecimal `json:"fp_jx_lxfy_12m"`               //
	SwCwbbSszb              decimal.NullDecimal `json:"sw_cwbb_sszb"`                 //
	FpJySychjeZb12mLh       decimal.NullDecimal `json:"fp_jy_sychje_zb_12m_lh"`       //
	FpJxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_jx_zyjyjezb_12m_lh"`        //
	FpXxXychjeZb12mLh       decimal.NullDecimal `json:"fp_xx_xychje_zb_12m_lh"`       //
	FpXxZyjyjezb12mLh       decimal.NullDecimal `json:"fp_xx_zyjyjezb_12m_lh"`        //
	SwSbQbxse12m            decimal.NullDecimal `json:"sw_sb_qbxse_12m"`              //
	SwSbQbxsezzl12m         decimal.NullDecimal `json:"sw_sb_qbxsezzl_12m"`           //
	SwSbLsxs12m             decimal.NullDecimal `json:"sw_sb_lsxs_12m"`               //
	SwCwbbChzztsCb          decimal.NullDecimal `json:"sw_cwbb_chzzts_cb"`            //
	SwCwbbZcfzl             decimal.NullDecimal `json:"sw_cwbb_zcfzl"`                //
	SwCwbbMlrzzlv           decimal.NullDecimal `json:"sw_cwbb_mlrzzlv"`              //
	SwCwbbJlrzzlv           decimal.NullDecimal `json:"sw_cwbb_jlrzzlv"`              //
	SwCwbbJzcszlv           decimal.NullDecimal `json:"sw_cwbb_jzcszlv"`              //
	SwJcxxClnx              decimal.NullDecimal `json:"sw_jcxx_clnx"`                 //
	StatusCode              int                 `json:"-"`                            //
	DecisionInputParam
}

type DecisionInputParam struct {
	ApplyTime string `json:"apply_time"`
	OrderNo   string `json:"order_no"`
}

func (s *RcDecisionParamDecisionRequestBody) Assignment(model *models.RcDecisionParam, dip *DecisionInputParam) {
	s.Id = model.Id
	s.ContentId = model.ContentId
	s.SwSdsnbCyrs = model.SwSdsnbCyrs
	s.GsGdct = model.GsGdct
	s.GsGdwdx = model.GsGdwdx
	s.GsFrwdx = model.GsFrwdx
	s.LhCylwz = model.LhCylwz
	s.LhMdPpjzl = model.LhMdPpjzl
	s.MdQybq = model.MdQybq
	s.SwCwbbYyzjzzts = model.SwCwbbYyzjzzts
	s.SfFhSsqkQy = model.SfFhSsqkQy
	s.SwJcxxNsrxypj = model.SwJcxxNsrxypj
	s.ZxYhsxqk = model.ZxYhsxqk
	s.ZxDsfsxqk = model.ZxDsfsxqk
	s.LhQylx = model.LhQylx
	s.Nsrsbh = model.Nsrsbh
	s.SwSbNszeZzsqysds12m = model.SwSbNszeZzsqysds12m
	s.SwSbNszezzlZzsqysds12mA = model.SwSbNszezzlZzsqysds12mA
	s.SwSdsnbGzxjzzjezzl = model.SwSdsnbGzxjzzjezzl
	s.SwSbzsSflhypld12m = model.SwSbzsSflhypld12m
	s.SwSdsnbYjfy = model.SwSdsnbYjfy
	s.FpJxLxfy12m = model.FpJxLxfy12m
	s.SwCwbbSszb = model.SwCwbbSszb
	s.FpJySychjeZb12mLh = model.FpJySychjeZb12mLh
	s.FpJxZyjyjezb12mLh = model.FpJxZyjyjezb12mLh
	s.FpXxXychjeZb12mLh = model.FpXxXychjeZb12mLh
	s.FpXxZyjyjezb12mLh = model.FpXxZyjyjezb12mLh
	s.SwSbQbxse12m = model.SwSbQbxse12m
	s.SwSbQbxsezzl12m = model.SwSbQbxsezzl12m
	s.SwSbLsxs12m = model.SwSbLsxs12m
	s.SwCwbbChzztsCb = model.SwCwbbChzztsCb
	s.SwCwbbZcfzl = model.SwCwbbZcfzl
	s.SwCwbbMlrzzlv = model.SwCwbbMlrzzlv
	s.SwCwbbJlrzzlv = model.SwCwbbJlrzzlv
	s.SwCwbbJzcszlv = model.SwCwbbJzcszlv
	s.SwJcxxClnx = model.SwJcxxClnx
	s.StatusCode = model.StatusCode
	s.ApplyTime = dip.ApplyTime
	s.OrderNo = dip.OrderNo
}

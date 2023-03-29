package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/pkg/errors"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common"
	"go-admin/common/apis"
	"go-admin/common/jwtauth/user"
	"strconv"
)

const INDEX = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>GO-ADMIN欢迎您</title>
<style>
body{
  margin:0; 
  padding:0; 
  overflow-y:hidden
}
</style>
<script src="http://libs.baidu.com/jquery/1.9.0/jquery.js"></script>
<script type="text/javascript"> 
window.onerror=function(){return true;} 
$(function(){ 
  headerH = 0;  
  var h=$(window).height();
  $("#iframe").height((h-headerH)+"px"); 
});
</script>
</head>
<body>
<iframe id="iframe" frameborder="0" src="https://doc.go-admin.dev" style="width:100%;"></iframe>
</body>
</html>
`

type Dashboard struct {
	apis.Api
}

func (e Dashboard) GoAdmin(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, INDEX)
}

func (e Dashboard) DashboardA(c *gin.Context) {
	req := dto.SysRequestLogGetPvReq{}
	req2 := dto.SysRequestLogGetPvReq{}
	req3 := dto.SysRequestLogGetPvReq{}
	req4 := dto.SysRequestLogGetPvReq{}
	req5 := dto.SysRequestLogGetUriReq{}
	resp := make([]dto.SysRequestLogGetPvResp, 0)
	resp2 := make([]dto.SysRequestLogGetPvResp, 0)
	resp3 := dto.SysRequestLogGetPvWithTypeResp{}
	resp4 := dto.SysRequestLogGetPvWithTypeResp{}
	resp5 := make([]dto.SysRequestLogGetPvResp, 0)
	s := service.SysRequestLog{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	if req.BeginTime != nil {
		end := now.With(*req.BeginTime).BeginningOfDay()
		req.BeginTime = &end
	}
	if req.EndTime != nil {
		end := now.With(*req.EndTime).EndOfDay()
		req.EndTime = &end
	}
	err = s.GetPv(&req, &resp)
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	req2.Type = "month"
	err = s.GetPv(&req2, &resp2)
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}

	begin3 := now.BeginningOfDay()
	end3 := now.EndOfDay()
	req3.BeginTime = &begin3
	req3.EndTime = &end3
	err = s.GetPvWithType(&req3, &resp3)
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	begin4 := now.BeginningOfYear()
	end4 := now.EndOfYear()
	req4.BeginTime = &begin4
	req4.EndTime = &end4
	err = s.GetPvWithType(&req4, &resp4)
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	req5.BeginTime = &begin4
	req5.EndTime = &end4
	err = s.GetMethod(&req5, &resp5)
	if err != nil {
		e.Logger.Error(err)
		panic(err)
		return
	}
	var str = T{}
	str.DayPv = resp3.Y
	str.TotalPv = resp4.Y
	if len(resp2) > 0 {
		month := now.BeginningOfMonth()
		monthEnd := now.EndOfMonth()
		for i := month; i.Unix() <= monthEnd.Unix(); i = i.AddDate(0, 0, 1) {
			xy := XY{X: i.Format("02"), Y: 0}
			for _, obj := range resp2 {
				t, err := now.Parse(obj.X)
				if err == nil && t.Unix() == i.Unix() {
					xy.Y = obj.Y
				}
			}
			str.VisitData = append(str.VisitData, xy)
		}
	}
	//str.VisitData = append(str.VisitData,
	//	XY{"2021-08-06", 7},
	//	XY{"2021-08-07", 5},
	//	XY{"2021-08-08", 4},
	//	XY{"2021-08-09", 2},
	//	XY{"2021-08-10", 4},
	//	XY{"2021-08-11", 7},
	//	XY{"2021-08-12", 5},
	//	XY{"2021-08-13", 6},
	//	XY{"2021-08-14", 5},
	//	XY{"2021-08-15", 9},
	//	XY{"2021-08-16", 6},
	//	XY{"2021-08-17", 3},
	//	XY{"2021-08-18", 1},
	//	XY{"2021-08-19", 5},
	//	XY{"2021-08-20", 3},
	//	XY{"2021-08-21", 6},
	//	XY{"2021-08-22", 5})
	str.VisitData2 = append(str.VisitData2,
		XY{"2021-08-06", 1},
		XY{"2021-08-07", 6},
		XY{"2021-08-08", 4},
		XY{"2021-08-09", 8},
		XY{"2021-08-10", 3},
		XY{"2021-08-11", 7},
		XY{"2021-08-12", 2})
	if req.Type == "month" {
		month := now.BeginningOfMonth()
		monthEnd := now.EndOfMonth()
		for i := month; i.Unix() <= monthEnd.Unix(); i = i.AddDate(0, 0, 1) {
			xy := XY{X: i.Format("02"), Y: 0}
			for _, obj := range resp {
				t, err := now.Parse(obj.X)
				if err == nil && t.Unix() == i.Unix() {
					xy.Y = obj.Y
				}
			}
			str.SalesData = append(str.SalesData, xy)
		}
	} else if req.Type == "" && req.BeginTime != nil && req.EndTime != nil && common.GetDiffDays(now.With(*req.EndTime).EndOfDay(), now.With(*req.BeginTime).BeginningOfDay()) != 0 {
		//cc := common.GetDiffDays(*req.EndTime, *req.BeginTime)
		//fmt.Println(cc)
		month := now.With(*req.BeginTime).Time
		if err != nil {
			err = errors.WithStack(err)
			e.Logger.Error(err)
		}

		monthEnd := now.With(*req.EndTime)
		if err != nil {
			err = errors.WithStack(err)
			e.Logger.Error(err)
		}
		for i := month; i.Unix() <= monthEnd.Unix(); i = i.AddDate(0, 0, 1) {
			xy := XY{X: i.Format("2006-01-02"), Y: 0}
			for _, obj := range resp {
				t, err := now.Parse(obj.X)
				if err == nil && t.Unix() == i.Unix() {
					xy.Y = obj.Y
				}
			}
			str.SalesData = append(str.SalesData, xy)
		}
	} else if req.Type == "week" {
		week := now.BeginningOfWeek()
		for i := 1; i < 8; i++ {
			xy := XY{X: week.Weekday().String(), Y: 0}
			for _, obj := range resp {
				if err == nil && obj.X == week.Format("2006-01-02") {
					xy.Y = obj.Y
				}
			}
			str.SalesData = append(str.SalesData, xy)
			week = week.AddDate(0, 0, 1)
		}
	} else if req.Type == "today" || (req.BeginTime != nil && req.EndTime != nil && common.GetDiffDays(*req.EndTime, *req.BeginTime) == 0) {
		for i := 0; i <= 24; i++ {
			xy := XY{strconv.Itoa(i) + "时", 0}
			for _, obj := range resp {
				m, err := strconv.Atoi(obj.X)
				if err == nil && m == i {
					xy.Y = obj.Y
				}
			}
			str.SalesData = append(str.SalesData, xy)
		}
	} else {
		for i := 1; i < 13; i++ {
			xy := XY{strconv.Itoa(i) + "月", 0}
			for _, obj := range resp {
				m, err := strconv.Atoi(obj.X)
				if err == nil && m == i {
					xy.Y = obj.Y
				}
			}
			str.SalesData = append(str.SalesData, xy)
		}
	}

	//str.SalesData = append(str.SalesData,
	//	XY{"1月", 752},
	//	XY{"2月", 598},
	//	XY{"3月", 890},
	//	XY{"4月", 285},
	//	XY{"5月", 691},
	//	XY{"6月", 800},
	//	XY{"7月", 887},
	//	XY{"8月", 1109},
	//	XY{"9月", 825},
	//	XY{"10月", 412},
	//	XY{"11月", 673},
	//	XY{"12月", 816})
	str.SearchData = append(str.SearchData,
		SearchData{1, "搜索关键词-0", 187, 57, 0},
		SearchData{2, "搜索关键词-1", 66, 70, 1},
		SearchData{3, "搜索关键词-2", 7, 1, 1},
		SearchData{4, "搜索关键词-3", 90, 93, 1},
		SearchData{5, "搜索关键词-4", 41, 49, 1},
		SearchData{6, "搜索关键词-5", 81, 16, 1},
		SearchData{7, "搜索关键词-6", 24, 41, 0},
		SearchData{8, "搜索关键词-7", 42, 47, 0},
		SearchData{9, "搜索关键词-8", 46, 26, 1},
		SearchData{10, "搜索关键词-9", 64, 84, 0},
		SearchData{11, "搜索关键词-10", 539, 54, 0},
		SearchData{12, "搜索关键词-11", 807, 5, 0},
		SearchData{13, "搜索关键词-12", 164, 14, 0},
		SearchData{14, "搜索关键词-13", 574, 12, 1},
		SearchData{15, "搜索关键词-14", 222, 24, 1},
		SearchData{16, "搜索关键词-15", 459, 2, 1},
		SearchData{17, "搜索关键词-16", 379, 83, 1},
		SearchData{18, "搜索关键词-17", 320, 47, 1},
		SearchData{19, "搜索关键词-18", 520, 96, 0},
		SearchData{20, "搜索关键词-19", 54, 30, 0},
		SearchData{21, "搜索关键词-20", 225, 7, 0},
		SearchData{22, "搜索关键词-21", 129, 91, 0},
		SearchData{23, "搜索关键词-22", 282, 16, 1},
		SearchData{24, "搜索关键词-23", 855, 44, 0},
		SearchData{25, "搜索关键词-24", 99, 21, 1},
		SearchData{26, "搜索关键词-25", 927, 7, 1},
		SearchData{27, "搜索关键词-26", 276, 68, 0},
		SearchData{28, "搜索关键词-27", 358, 65, 0},
		SearchData{29, "搜索关键词-28", 527, 29, 0},
		SearchData{30, "搜索关键词-29", 516, 34, 0},
		SearchData{31, "搜索关键词-30", 594, 55, 1},
		SearchData{32, "搜索关键词-31", 721, 69, 0},
		SearchData{33, "搜索关键词-32", 957, 58, 0},
		SearchData{34, "搜索关键词-33", 13, 48, 0},
		SearchData{35, "搜索关键词-34", 251, 69, 0},
		SearchData{36, "搜索关键词-35", 603, 86, 0},
		SearchData{37, "搜索关键词-36", 459, 47, 1},
		SearchData{38, "搜索关键词-37", 566, 12, 1},
		SearchData{39, "搜索关键词-38", 285, 36, 0},
		SearchData{40, "搜索关键词-39", 899, 28, 0},
		SearchData{41, "搜索关键词-40", 201, 63, 0},
		SearchData{42, "搜索关键词-41", 6, 46, 1},
		SearchData{43, "搜索关键词-42", 587, 87, 0},
		SearchData{44, "搜索关键词-43", 58, 20, 1},
		SearchData{45, "搜索关键词-44", 521, 61, 1},
		SearchData{46, "搜索关键词-45", 112, 74, 1},
		SearchData{47, "搜索关键词-46", 402, 5, 0},
		SearchData{48, "搜索关键词-47", 101, 54, 0},
		SearchData{49, "搜索关键词-48", 918, 58, 0},
		SearchData{50, "搜索关键词-49", 322, 75, 0})
	str.OfflineData = append(str.OfflineData,
		OfflineData{"Stores 0", 0.4},
		OfflineData{"Stores 1", 0.3},
		OfflineData{"Stores 2", 0.1},
		OfflineData{"Stores 3", 0.2},
		OfflineData{"Stores 4", 0.5},
		OfflineData{"Stores 5", 0.1},
		OfflineData{"Stores 6", 0.9},
		OfflineData{"Stores 7", 0.6},
		OfflineData{"Stores 8", 0.7},
		OfflineData{"Stores 9", 0.7})
	str.OfflineChartData = append(str.OfflineChartData,
		OfflineChartData{"11:32", "客流量", 81},
		OfflineChartData{"11:32", "支付笔数", 21},
		OfflineChartData{"12:02", "客流量", 37},
		OfflineChartData{"12:02", "支付笔数", 31},
		OfflineChartData{"12:32", "客流量", 56},
		OfflineChartData{"12:32", "支付笔数", 87},
		OfflineChartData{"13:02", "客流量", 26},
		OfflineChartData{"13:02", "支付笔数", 38},
		OfflineChartData{"13:32", "客流量", 70},
		OfflineChartData{"13:32", "支付笔数", 64},
		OfflineChartData{"14:02", "客流量", 96},
		OfflineChartData{"14:02", "支付笔数", 94},
		OfflineChartData{"14:32", "客流量", 103},
		OfflineChartData{"14:32", "支付笔数", 39},
		OfflineChartData{"15:02", "客流量", 50},
		OfflineChartData{"15:02", "支付笔数", 105},
		OfflineChartData{"15:32", "客流量", 75},
		OfflineChartData{"15:32", "支付笔数", 80},
		OfflineChartData{"16:02", "客流量", 100},
		OfflineChartData{"16:02", "支付笔数", 13},
		OfflineChartData{"16:32", "客流量", 67},
		OfflineChartData{"16:32", "支付笔数", 68},
		OfflineChartData{"17:02", "客流量", 90},
		OfflineChartData{"17:02", "支付笔数", 45},
		OfflineChartData{"17:32", "客流量", 31},
		OfflineChartData{"17:32", "支付笔数", 21},
		OfflineChartData{"18:02", "客流量", 58},
		OfflineChartData{"18:02", "支付笔数", 64},
		OfflineChartData{"18:32", "客流量", 10},
		OfflineChartData{"18:32", "支付笔数", 64},
		OfflineChartData{"19:02", "客流量", 79},
		OfflineChartData{"19:02", "支付笔数", 18})
	//str.SalesTypeData = append(str.SalesTypeData,
	//	XY{"家用电器", 4544})
	for _, obj := range resp5 {
		//if obj.X != "GET" {
		str.SalesTypeData = append(str.SalesTypeData, XY{obj.X, obj.Y})
		str.SalesTypeDataOnline = append(str.SalesTypeDataOnline, XY{obj.X, obj.Y})
		str.SalesTypeDataOffline = append(str.SalesTypeDataOffline, XY{obj.X, obj.Y})
		//}
	}
	//str.SalesTypeData = append(str.SalesTypeData,
	//	XY{"家用电器", 4544},
	//	XY{"食用酒水", 3321},
	//	XY{"个护健康", 3113},
	//	XY{"服饰箱包", 2341},
	//	XY{"母婴产品", 1231},
	//	XY{"其他", 1231})
	//str.SalesTypeDataOnline = append(str.SalesTypeDataOnline,
	//	XY{"家用电器", 244},
	//	XY{"食用酒水", 321},
	//	XY{"个护健康", 311},
	//	XY{"服饰箱包", 41},
	//	XY{"母婴产品", 121},
	//	XY{"其他", 111})
	//str.SalesTypeDataOffline = append(str.SalesTypeDataOffline,
	//	XY{"家用电器", 99},
	//	XY{"食用酒水", 188},
	//	XY{"个护健康", 344},
	//	XY{"服饰箱包", 255},
	//	XY{"其他", 65})
	str.RadarData = append(str.RadarData,
		RadarData{"个人", "引用", 10},
		RadarData{"个人", "口碑", 8},
		RadarData{"个人", "产量", 4},
		RadarData{"个人", "贡献", 5},
		RadarData{"个人", "热度", 7},
		RadarData{"团队", "引用", 3},
		RadarData{"团队", "口碑", 9},
		RadarData{"团队", "产量", 6},
		RadarData{"团队", "贡献", 3},
		RadarData{"团队", "热度", 1},
		RadarData{"部门", "引用", 4},
		RadarData{"部门", "口碑", 1},
		RadarData{"部门", "产量", 6},
		RadarData{"部门", "贡献", 5},
		RadarData{"部门", "热度", 7})

	e.OK(str)
}

type XY struct {
	X string `json:"x"`
	Y int    `json:"y"`
}
type RadarData struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value int    `json:"value"`
}
type OfflineChartData struct {
	Date  string `json:"date"`
	Type  string `json:"type"`
	Value int    `json:"value"`
}
type SearchData struct {
	Index   int    `json:"index"`
	Keyword string `json:"keyword"`
	Count   int    `json:"count"`
	Range   int    `json:"range"`
	Status  int    `json:"status"`
}
type OfflineData struct {
	Name string  `json:"name"`
	Cvr  float64 `json:"cvr"`
}
type T struct {
	VisitData            []XY               `json:"visitData"`
	VisitData2           []XY               `json:"visitData2"`
	SalesData            []XY               `json:"salesData"`
	SearchData           []SearchData       `json:"searchData"`
	OfflineData          []OfflineData      `json:"offlineData"`
	OfflineChartData     []OfflineChartData `json:"offlineChartData"`
	SalesTypeData        []XY               `json:"salesTypeData"`
	SalesTypeDataOnline  []XY               `json:"salesTypeDataOnline"`
	SalesTypeDataOffline []XY               `json:"salesTypeDataOffline"`
	RadarData            []RadarData        `json:"radarData"`
	TotalPv              int                `json:"totalPv"`
	DayPv                int                `json:"dayPv"`
}

func (e Dashboard) Activities(c *gin.Context) {
	e.MakeContext(c)
	var t2 = make([]T2, 0)
	t2 = append(t2, T2{Id: "trend-1", UpdatedAt: "2021-08-12T09:19:58.845Z",
		User:     User{Name: "曲丽丽", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png"},
		Group:    Group{Name: "高逼格设计天团", Link: "http://github.com/"},
		Project:  Project{Name: "六月迭代", Link: "http://github.com/"},
		Template: "在 @{group} 新建项目 @{project}",
	},
		T2{Id: "trend-2", UpdatedAt: "2021-08-12T09:19:58.845Z",
			User:     User{Name: "付小小", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/cnrhVkzwxjPwAaCfPbdc.png"},
			Group:    Group{Name: "高逼格设计天团", Link: "http://github.com/"},
			Project:  Project{Name: "六月迭代", Link: "http://github.com/"},
			Template: "在 @{group} 新建项目 @{project}",
		},
		T2{Id: "trend-3", UpdatedAt: "2021-08-12T09:19:58.845Z",
			User:     User{Name: "林东东", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/gaOngJwsRYRaVAuXXcmB.png"},
			Group:    Group{Name: "中二少女团", Link: "http://github.com/"},
			Project:  Project{Name: "六月迭代", Link: "http://github.com/"},
			Template: "在 @{group} 新建项目 @{project}",
		},
		T2{Id: "trend-4", UpdatedAt: "2021-08-12T09:19:58.845Z",
			User:     User{Name: "周星星", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/WhxKECPNujWoWEFNdnJE.png"},
			Project:  Project{Name: "5 月日常迭代", Link: "http://github.com/"},
			Template: "将 @{project} 更新至已发布状态",
		},
		T2{Id: "trend-5", UpdatedAt: "2021-08-12T09:19:58.845Z",
			User:     User{Name: "朱偏右", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/ubnKSIfAJTxIgXOKlciN.png"},
			Project:  Project{Name: "工程效能", Link: "http://github.com/"},
			Comment:  Comment{Name: "留言", Link: "http://github.com/"},
			Template: "在 @{project} 发布了 @{comment}",
		},
		T2{Id: "trend-6", UpdatedAt: "2021-08-12T09:19:58.845Z",
			User:     User{"乐哥", "https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png"},
			Group:    Group{"程序员日常", "http://github.com/"},
			Project:  Project{"品牌迭代", "http://github.com/"},
			Template: "在 @{group} 新建项目 @{project}",
		})
	e.OK(t2)
}

type User struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
type Group struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
type Project struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
type Comment struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type T2 struct {
	Id        string  `json:"id"`
	UpdatedAt string  `json:"updatedAt"`
	User      User    `json:"user"`
	Group     Group   `json:"group,omitempty"`
	Project   Project `json:"project"`
	Template  string  `json:"template"`
	Comment   Comment `json:"comment,omitempty"`
}

func (e Dashboard) Notice(c *gin.Context) {
	e.MakeContext(c)
	var t3 = make([]T3, 0)
	t3 = append(t3,
		T3{Id: "xxx1", Title: "Alipay", Logo: "https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png", Description: "那是一种内在的东西，他们到达不了，也无法触及的", UpdatedAt: "2021-08-12T09:19:58.815Z", Member: "科学搬砖组", Href: "", MemberLink: ""},
		T3{Id: "xxx2", Title: "Angular", Logo: "https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png", Description: "希望是一个好东西，也许是最好的，好东西是不会消亡的", UpdatedAt: "2017-07-24T00:00:00.000Z", Member: "全组都是吴彦祖", Href: "", MemberLink: ""},
		T3{Id: "xxx3", Title: "Ant Design", Logo: "https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png", Description: "城镇中有那么多的酒馆，她却偏偏走进了我的酒馆", UpdatedAt: "2021-08-12T09:19:58.815Z", Member: "中二少女团", Href: "", MemberLink: ""},
		T3{Id: "xxx4", Title: "Ant Design Pro", Logo: "https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png", Description: "那时候我只会想自己想要什么，从不想自己拥有什么", UpdatedAt: "2017-07-23T00:00:00.000Z", Member: "程序员日常", Href: "", MemberLink: ""},
		T3{Id: "xxx5", Title: "Bootstrap", Logo: "https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png", Description: "凛冬将至", UpdatedAt: "2017-07-23T00:00:00.000Z", Member: "高逼格设计天团", Href: "", MemberLink: ""},
		T3{Id: "xxx6", Title: "React", Logo: "https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png", Description: "生命就像一盒巧克力，结果往往出人意料", UpdatedAt: "2017-07-23T00:00:00.000Z", Member: "骗你来学计算机", Href: "", MemberLink: ""},
	)
	e.OK(t3)
}

type T3 struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updatedAt"`
	Member      string `json:"member"`
	Href        string `json:"href"`
	MemberLink  string `json:"memberLink"`
}

func (e Dashboard) Notices(c *gin.Context) {
	e.MakeContext(c)
	t4 := make([]T4, 0)
	t4 = append(t4,
		T4{Id: "000000001", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png", Title: "你收到了 14 份新周报", Datetime: "2017-08-09", Type: "notification"},
		T4{Id: "000000002", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/OKJXDXrmkNshAMvwtvhu.png", Title: "你推荐的 曲妮妮 已通过第三轮面试", Datetime: "2017-08-08", Type: "notification"},
		T4{Id: "000000003", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/kISTdvpyTAhtGxpovNWd.png", Title: "这种模板可以区分多种通知类型", Datetime: "2017-08-07", Read: true, Type: "notification"},
		T4{Id: "000000004", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/GvqBnKhFgObvnSGkDsje.png", Title: "左侧图标用于区分不同的类型", Datetime: "2017-08-07", Type: "notification"},
		T4{Id: "000000005", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png", Title: "内容不要超过两行字，超出时自动截断", Datetime: "2017-08-07", Type: "notification"},
		T4{Id: "000000006", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/fcHMVNCjPOsbUGdEduuv.jpeg", Title: "曲丽丽 评论了你", Description: "描述信息描述信息描述信息", Datetime: "2017-08-07", Type: "message", ClickClose: true},
		T4{Id: "000000007", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/fcHMVNCjPOsbUGdEduuv.jpeg", Title: "朱偏右 回复了你", Description: "这种模板用于提醒谁与你发生了互动，左侧放『谁』的头像", Datetime: "2017-08-07", Type: "message", ClickClose: true},
		T4{Id: "000000008", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/fcHMVNCjPOsbUGdEduuv.jpeg", Title: "标题", Description: "这种模板用于提醒谁与你发生了互动，左侧放『谁』的头像", Datetime: "2017-08-07", Type: "message", ClickClose: true},
		T4{Id: "000000009", Title: "任务名称", Description: "任务需要在 2017-01-12 20:00 前启动", Extra: "未开始", Status: "todo", Type: "event"},
		T4{Id: "000000010", Title: "第三方紧急代码变更", Description: "冠霖提交于 2017-01-06，需在 2017-01-07 前完成代码变更任务", Extra: "马上到期", Status: "urgent", Type: "event"},
		T4{Id: "000000011", Title: "信息安全考试", Description: "指派竹尔于 2017-01-09 前完成更新并发布", Extra: "已耗时 8 天", Status: "doing", Type: "event"},
		T4{Id: "000000012", Title: "ABCD 版本发布", Description: "冠霖提交于 2017-01-06，需在 2017-01-07 前完成代码变更任务", Extra: "进行中", Status: "processing", Type: "event"},
	)
	e.OK(t4)
}

type T4 struct {
	Id          string `json:"id"`
	Avatar      string `json:"avatar,omitempty"`
	Title       string `json:"title"`
	Datetime    string `json:"datetime,omitempty"`
	Type        string `json:"type"`
	Read        bool   `json:"read,omitempty"`
	Description string `json:"description,omitempty"`
	ClickClose  bool   `json:"clickClose,omitempty"`
	Extra       string `json:"extra,omitempty"`
	Status      string `json:"status,omitempty"`
}

func (e Dashboard) CurrentUserDetail(c *gin.Context) {
	e.MakeContext(c)
	userName := user.GetUserName(c)
	tags := make([]Tag, 0)
	tags = append(tags, Tag{"0", "很有想法的"},
		Tag{"1", "专注设计"},
		Tag{"2", "辣~"},
		Tag{"3", "大长腿"},
		Tag{"4", "川妹子"},
		Tag{"5", "海纳百川"})
	notices := make([]Notice, 0)
	notices = append(notices,
		Notice{Id: "xxx1", Title: "Alipay", Logo: "https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png", Description: "那是一种内在的东西，他们到达不了，也无法触及的", UpdatedAt: "2021-08-11T04:22:58.510Z", Member: "科学搬砖组", Href: "", MemberLink: ""},
		Notice{Id: "xxx2", Title: "Angular", Logo: "https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png", Description: "希望是一个好东西，也许是最好的，好东西是不会消亡的", UpdatedAt: "2017-07-24T00:00:00.000Z", Member: "全组都是吴彦祖", Href: "", MemberLink: ""},
		Notice{Id: "xxx3", Title: "Ant Design", Logo: "https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png", Description: "城镇中有那么多的酒馆，她却偏偏走进了我的酒馆", UpdatedAt: "2021-08-11T04:22:58.510Z", Member: "中二少女团", Href: "", MemberLink: ""},
		Notice{Id: "xxx4", Title: "Ant Design Pro", Logo: "https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png", Description: "那时候我只会想自己想要什么，从不想自己拥有什么", UpdatedAt: "2017-07-23T00:00:00.000Z", Member: "程序员日常", Href: "", MemberLink: ""},
		Notice{Id: "xxx5", Title: "Bootstrap", Logo: "https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png", Description: "凛冬将至", UpdatedAt: "2017-07-23T00:00:00.000Z", Member: "高逼格设计天团", Href: "", MemberLink: ""},
		Notice{Id: "xxx6", Title: "React", Logo: "https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png", Description: "生命就像一盒巧克力，结果往往出人意料", UpdatedAt: "2017-07-23T00:00:00.000Z", Member: "骗你来学计算机", Href: "", MemberLink: ""},
	)
	t5 := T5{Name: userName,
		Avatar:      "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		Userid:      "00000001",
		Email:       "antdesign@alipay.com",
		Signature:   "海纳百川，有容乃大",
		Title:       "交互专家",
		Group:       "蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED",
		Tags:        tags,
		Notice:      notices,
		NotifyCount: 12, UnreadCount: 11, Country: "China",
		Geographic: Geographic{
			Province: KV{Label: "浙江省", Key: "330000"},
			City:     KV{Label: "杭州市", Key: "330100"}},
		Address: "西湖区工专路 77 号",
		Phone:   "0752-268888888"}

	e.OK(t5)
}

type Tag struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}
type Notice struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updatedAt"`
	Member      string `json:"member"`
	Href        string `json:"href"`
	MemberLink  string `json:"memberLink"`
}
type KV struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}
type Geographic struct {
	Province KV `json:"province"`
	City     KV `json:"city"`
}
type T5 struct {
	Name        string     `json:"name"`
	Avatar      string     `json:"avatar"`
	Userid      string     `json:"userid"`
	Email       string     `json:"email"`
	Signature   string     `json:"signature"`
	Title       string     `json:"title"`
	Group       string     `json:"group"`
	Tags        []Tag      `json:"tags"`
	Notice      []Notice   `json:"notice"`
	NotifyCount int        `json:"notifyCount"`
	UnreadCount int        `json:"unreadCount"`
	Country     string     `json:"country"`
	Geographic  Geographic `json:"geographic"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
}

func (e Dashboard) FakeListDetail(c *gin.Context) {
	e.MakeContext(c)
	members := make([]Member, 0)
	members = append(members,
		Member{Avatar: "https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png", Name: "曲丽丽", Id: "member1"},
		Member{Avatar: "https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png", Name: "王昭君", Id: "member2"},
		Member{Avatar: "https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png", Name: "董娜娜", Id: "member3"})
	t6 := make([]T6, 0)
	t6 = append(t6,
		T6{Id: "fake-list-0", Owner: "付小小", Title: "Alipay", Avatar: "https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png", Cover: "https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png", Status: "active", Percent: 88, Logo: "https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png", Href: "https://ant.design", UpdatedAt: 1628765168350, CreatedAt: 1628765168350, SubDescription: "那是一种内在的东西， 他们到达不了，也无法触及的", Description: "在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。", ActiveUser: 165660, NewUser: 1233, Star: 177, Like: 139, Message: 18, Content: "段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
			Members: members,
		})
	//{"id":"fake-list-1","owner":"曲丽丽","title":"Angular","avatar":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png","status":"exception","percent":84,"logo":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","href":"https://ant.design","updatedAt":1628757968350,"createdAt":1628757968350,"subDescription":"希望是一个好东西，也许是最好的，好东西是不会消亡的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":101516,"newUser":1702,"star":193,"like":119,"message":19,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-2","owner":"林东东","title":"Ant Design","avatar":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png","status":"normal","percent":64,"logo":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","href":"https://ant.design","updatedAt":1628750768350,"createdAt":1628750768350,"subDescription":"生命就像一盒巧克力，结果往往出人意料","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":191910,"newUser":1473,"star":121,"like":163,"message":18,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-3","owner":"周星星","title":"Ant Design Pro","avatar":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png","status":"active","percent":83,"logo":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","href":"https://ant.design","updatedAt":1628743568350,"createdAt":1628743568350,"subDescription":"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":106249,"newUser":1427,"star":129,"like":101,"message":12,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-4","owner":"吴加好","title":"Bootstrap","avatar":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png","status":"exception","percent":70,"logo":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","href":"https://ant.design","updatedAt":1628736368350,"createdAt":1628736368350,"subDescription":"那时候我只会想自己想要什么，从不想自己拥有什么","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":120929,"newUser":1312,"star":186,"like":191,"message":12,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]},
	//{"id":"fake-list-5","owner":"朱偏右","title":"React","avatar":"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png","status":"normal","percent":84,"logo":"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png","href":"https://ant.design","updatedAt":1628729168350,"createdAt":1628729168350,"subDescription":"那是一种内在的东西， 他们到达不了，也无法触及的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":146509,"newUser":1972,"star":161,"like":129,"message":12,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-6","owner":"鱼酱","title":"Vue","avatar":"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png","status":"active","percent":64,"logo":"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png","href":"https://ant.design","updatedAt":1628721968350,"createdAt":1628721968350,"subDescription":"希望是一个好东西，也许是最好的，好东西是不会消亡的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":144822,"newUser":1486,"star":174,"like":104,"message":15,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]},
	//{"id":"fake-list-7","owner":"乐哥","title":"Webpack","avatar":"https://gw.alipayobjects.com/zos/rmsportal/nxkuOJlFJuAUhzlMTCEe.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png","status":"exception","percent":73,"logo":"https://gw.alipayobjects.com/zos/rmsportal/nxkuOJlFJuAUhzlMTCEe.png","href":"https://ant.design","updatedAt":1628714768350,"createdAt":1628714768350,"subDescription":"生命就像一盒巧克力，结果往往出人意料","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":152410,"newUser":1328,"star":148,"like":195,"message":20,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-8","owner":"谭小仪","title":"Alipay","avatar":"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png","status":"normal","percent":65,"logo":"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png","href":"https://ant.design","updatedAt":1628707568350,"createdAt":1628707568350,"subDescription":"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":184535,"newUser":1452,"star":170,"like":134,"message":14,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-9","owner":"仲尼","title":"Angular","avatar":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png","status":"active","percent":99,"logo":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","href":"https://ant.design","updatedAt":1628700368350,"createdAt":1628700368350,"subDescription":"那时候我只会想自己想要什么，从不想自己拥有什么","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":167233,"newUser":1421,"star":190,"like":178,"message":17,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-10","owner":"付小小","title":"Ant Design","avatar":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png","status":"exception","percent":92,"logo":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","href":"https://ant.design","updatedAt":1628693168350,"createdAt":1628693168350,"subDescription":"那是一种内在的东西， 他们到达不了，也无法触及的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":102216,"newUser":1953,"star":166,"like":180,"message":13,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-11","owner":"曲丽丽","title":"Ant Design Pro","avatar":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png","status":"normal","percent":81,"logo":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","href":"https://ant.design","updatedAt":1628685968350,"createdAt":1628685968350,"subDescription":"希望是一个好东西，也许是最好的，好东西是不会消亡的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":180229,"newUser":1708,"star":146,"like":154,"message":13,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-12","owner":"林东东","title":"Bootstrap","avatar":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png","status":"active","percent":75,"logo":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","href":"https://ant.design","updatedAt":1628678768350,"createdAt":1628678768350,"subDescription":"生命就像一盒巧克力，结果往往出人意料","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":146921,"newUser":1704,"star":114,"like":110,"message":16,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-13","owner":"周星星","title":"React","avatar":"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png","status":"exception","percent":85,"logo":"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png","href":"https://ant.design","updatedAt":1628671568350,"createdAt":1628671568350,"subDescription":"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":175042,"newUser":1443,"star":119,"like":194,"message":14,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-14","owner":"吴加好","title":"Vue","avatar":"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png","status":"normal","percent":85,"logo":"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png","href":"https://ant.design","updatedAt":1628664368350,"createdAt":1628664368350,"subDescription":"那时候我只会想自己想要什么，从不想自己拥有什么","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":140092,"newUser":1429,"star":142,"like":184,"message":20,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-15","owner":"朱偏右","title":"Webpack","avatar":"https://gw.alipayobjects.com/zos/rmsportal/nxkuOJlFJuAUhzlMTCEe.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png","status":"active","percent":64,"logo":"https://gw.alipayobjects.com/zos/rmsportal/nxkuOJlFJuAUhzlMTCEe.png","href":"https://ant.design","updatedAt":1628657168350,"createdAt":1628657168350,"subDescription":"那是一种内在的东西， 他们到达不了，也无法触及的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":170343,"newUser":1491,"star":198,"like":158,"message":19,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-16","owner":"鱼酱","title":"Alipay","avatar":"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png","status":"exception","percent":69,"logo":"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png","href":"https://ant.design","updatedAt":1628649968350,"createdAt":1628649968350,"subDescription":"希望是一个好东西，也许是最好的，好东西是不会消亡的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":194154,"newUser":1058,"star":188,"like":177,"message":14,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-17","owner":"乐哥","title":"Angular","avatar":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png","status":"normal","percent":60,"logo":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","href":"https://ant.design","updatedAt":1628642768350,"createdAt":1628642768350,"subDescription":"生命就像一盒巧克力，结果往往出人意料","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":103728,"newUser":1854,"star":113,"like":196,"message":16,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-18","owner":"谭小仪","title":"Ant Design","avatar":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png","status":"active","percent":55,"logo":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","href":"https://ant.design","updatedAt":1628635568350,"createdAt":1628635568350,"subDescription":"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":175272,"newUser":1794,"star":176,"like":169,"message":16,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-19","owner":"仲尼","title":"Ant Design Pro","avatar":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png","status":"exception","percent":98,"logo":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","href":"https://ant.design","updatedAt":1628628368350,"createdAt":1628628368350,"subDescription":"那时候我只会想自己想要什么，从不想自己拥有什么","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":157046,"newUser":1351,"star":123,"like":196,"message":11,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-20","owner":"付小小","title":"Bootstrap","avatar":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png","status":"normal","percent":56,"logo":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","href":"https://ant.design","updatedAt":1628621168350,"createdAt":1628621168350,"subDescription":"那是一种内在的东西， 他们到达不了，也无法触及的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":186315,"newUser":1232,"star":103,"like":186,"message":18,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-21","owner":"曲丽丽","title":"React","avatar":"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png","status":"active","percent":76,"logo":"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png","href":"https://ant.design","updatedAt":1628613968350,"createdAt":1628613968350,"subDescription":"希望是一个好东西，也许是最好的，好东西是不会消亡的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":191898,"newUser":1313,"star":131,"like":125,"message":11,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-22","owner":"林东东","title":"Vue","avatar":"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png","status":"exception","percent":70,"logo":"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png","href":"https://ant.design","updatedAt":1628606768350,"createdAt":1628606768350,"subDescription":"生命就像一盒巧克力，结果往往出人意料","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":134563,"newUser":1012,"star":187,"like":192,"message":13,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-23","owner":"周星星","title":"Webpack","avatar":"https://gw.alipayobjects.com/zos/rmsportal/nxkuOJlFJuAUhzlMTCEe.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png","status":"normal","percent":96,"logo":"https://gw.alipayobjects.com/zos/rmsportal/nxkuOJlFJuAUhzlMTCEe.png","href":"https://ant.design","updatedAt":1628599568350,"createdAt":1628599568350,"subDescription":"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":172914,"newUser":1539,"star":172,"like":116,"message":19,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-24","owner":"吴加好","title":"Alipay","avatar":"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png","status":"active","percent":88,"logo":"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png","href":"https://ant.design","updatedAt":1628592368350,"createdAt":1628592368350,"subDescription":"那时候我只会想自己想要什么，从不想自己拥有什么","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":169352,"newUser":1405,"star":199,"like":144,"message":14,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-25","owner":"朱偏右","title":"Angular","avatar":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png","status":"exception","percent":70,"logo":"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png","href":"https://ant.design","updatedAt":1628585168350,"createdAt":1628585168350,"subDescription":"那是一种内在的东西， 他们到达不了，也无法触及的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":116057,"newUser":1218,"star":180,"like":103,"message":11,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-26","owner":"鱼酱","title":"Ant Design","avatar":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png","status":"normal","percent":70,"logo":"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png","href":"https://ant.design","updatedAt":1628577968350,"createdAt":1628577968350,"subDescription":"希望是一个好东西，也许是最好的，好东西是不会消亡的","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":150922,"newUser":1465,"star":192,"like":145,"message":15,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-27","owner":"乐哥","title":"Ant Design Pro","avatar":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png","status":"active","percent":56,"logo":"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png","href":"https://ant.design","updatedAt":1628570768350,"createdAt":1628570768350,"subDescription":"生命就像一盒巧克力，结果往往出人意料","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":189725,"newUser":1061,"star":200,"like":167,"message":15,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},
	//{"id":"fake-list-28","owner":"谭小仪","title":"Bootstrap","avatar":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png","status":"exception","percent":98,"logo":"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png","href":"https://ant.design","updatedAt":1628563568350,"createdAt":1628563568350,"subDescription":"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":181286,"newUser":1141,"star":115,"like":172,"message":16,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//	"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//},{"id":"fake-list-29","owner":"仲尼","title":"React","avatar":"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png","cover":"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png","status":"normal","percent":79,"logo":"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png","href":"https://ant.design","updatedAt":1628556368350,"createdAt":1628556368350,"subDescription":"那时候我只会想自己想要什么，从不想自己拥有什么","description":"在中台产品的研发过程中，会出现不同的设计规范和实现方式，但其中往往存在很多类似的页面和组件，这些类似的组件会被抽离成一套标准规范。","activeUser":145238,"newUser":1175,"star":108,"like":200,"message":20,"content":"段落示意：蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。蚂蚁金服设计平台 ant.design，用最小的工作量，无缝接入蚂蚁金服生态，提供跨越设计与开发的体验解决方案。",
	//		"members":[{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png","name":"曲丽丽","id":"member1"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png","name":"王昭君","id":"member2"},
	//	{"avatar":"https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png","name":"董娜娜","id":"member3"}]
	//}]}
	list := make(map[string]interface{}, 0)
	list["list"] = t6
	e.OK(list)
}

type Member struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Id     string `json:"id"`
}
type T6 struct {
	Id             string   `json:"id"`
	Owner          string   `json:"owner"`
	Title          string   `json:"title"`
	Avatar         string   `json:"avatar"`
	Cover          string   `json:"cover"`
	Status         string   `json:"status"`
	Percent        int      `json:"percent"`
	Logo           string   `json:"logo"`
	Href           string   `json:"href"`
	UpdatedAt      int64    `json:"updatedAt"`
	CreatedAt      int64    `json:"createdAt"`
	SubDescription string   `json:"subDescription"`
	Description    string   `json:"description"`
	ActiveUser     int      `json:"activeUser"`
	NewUser        int      `json:"newUser"`
	Star           int      `json:"star"`
	Like           int      `json:"like"`
	Message        int      `json:"message"`
	Content        string   `json:"content"`
	Members        []Member `json:"members"`
}

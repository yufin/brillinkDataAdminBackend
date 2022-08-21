package middleware

type UrlInfo struct {
	Url    string
	Method string
	Access string
}

// CasbinExclude 排除的路由列表
// 用于Casbin排除过滤的信息
var CasbinExclude = []UrlInfo{
	{Url: "/api/v1/user/current", Method: "GET"},
	{Url: "/api/v1/dict/type-option-select", Method: "GET"},
	{Url: "/api/v1/dict-data/option-select", Method: "GET"},
	{Url: "/api/v1/deptTree", Method: "GET"},
	{Url: "/api/v1/db/tables/page", Method: "GET"},
	{Url: "/api/v1/db/columns/page", Method: "GET"},
	{Url: "/api/v1/gen/toproject/:tableId", Method: "GET"},
	{Url: "/api/v1/gen/todb/:tableId", Method: "GET"},
	{Url: "/api/v1/gen/tabletree", Method: "GET"},
	{Url: "/api/v1/gen/preview/:tableId", Method: "GET"},
	{Url: "/api/v1/gen/apitofile/:tableId", Method: "GET"},
	{Url: "/api/v1/getCaptcha", Method: "GET"},
	{Url: "/api/v1/getinfo", Method: "GET"},
	{Url: "/api/v1/menuTreeselect", Method: "GET"},
	{Url: "/api/v1/menurole", Method: "GET"},
	{Url: "/api/v1/menuids", Method: "GET"},
	{Url: "/api/v1/roleMenuTreeselect/:roleId", Method: "GET"},
	{Url: "/api/v1/roleDeptTreeselect/:roleId", Method: "GET"},
	{Url: "/api/v1/refresh_token", Method: "GET"},
	{Url: "/api/v1/configKey/:configKey", Method: "GET"},
	{Url: "/api/v1/app-config", Method: "GET"},
	{Url: "/api/v1/user/profile", Method: "GET"},
	{Url: "/info", Method: "GET"},
	{Url: "/api/v1/login", Method: "POST"},
	{Url: "/api/v1/logout", Method: "POST"},
	{Url: "/api/v1/user/avatar", Method: "POST"},
	{Url: "/api/v1/user/pwd", Method: "PUT"},
	{Url: "/api/v1/metrics", Method: "GET"},
	{Url: "/api/v1/health", Method: "GET"},
	{Url: "/", Method: "GET"},
	{Url: "/api/v1/server-monitor", Method: "GET"},
	{Url: "/api/v1/public/uploadFile", Method: "POST"},
}

// RouterRule  业务路由配置
// 注意：map key 要和 菜单中的【权限标记】一模一样，否则会有问题！！！
var RouterRule = make(map[string][]UrlInfo, 3)

func init() {
	RouterRule["canDashboard"] = []UrlInfo{
		{Url: "api/v1/fake_workplace_chart_data", Method: "GET"},
	}
	RouterRule["canDashboardWorkplace"] = []UrlInfo{
		{Url: "/api/v1/fake_analysis_chart_data", Method: "GET"},
	}

	// canUserList
	RouterRule["canUserListGetPage"] = []UrlInfo{
		{Url: "/api/v1/sys-user", Method: "GET"},
	}
	RouterRule["canUserListNew"] = []UrlInfo{
		{Url: "/api/v1/sys-user", Method: "POST"},
	}
	RouterRule["canUserListEdit"] = []UrlInfo{
		{Url: "/api/v1/sys-user/:id", Method: "GET"},
		{Url: "/api/v1/sys-user", Method: "PUT"},
	}
	RouterRule["canUserListDelete"] = []UrlInfo{
		{Url: "/api/v1/sys-user", Method: "DELETE"},
	}
	RouterRule["canUserListResetPassword"] = []UrlInfo{
		{Url: "/api/v1/user/pwd/reset", Method: "PUT"},
	}

	// canMenuListGetPage
	RouterRule["canMenuListGetPage"] = []UrlInfo{
		{Url: "/api/v1/menu", Method: "GET"},
	}
	RouterRule["canMenuLisNew"] = []UrlInfo{
		{Url: "/api/v1/menu", Method: "POST"},
	}
	RouterRule["canMenuLisEdit"] = []UrlInfo{
		{Url: "/api/v1/menu/:id", Method: "GET"},
		{Url: "/api/v1/menu/:id", Method: "PUT"},
	}
	RouterRule["canMenuLisDelete"] = []UrlInfo{
		{Url: "/api/v1/menu", Method: "DELETE"},
	}

	// canRoleListGetPage
	RouterRule["canRoleListGetPage"] = []UrlInfo{
		{Url: "/api/v1/role", Method: "GET"},
	}
	RouterRule["canRoleListNew"] = []UrlInfo{
		{Url: "/api/v1/role", Method: "POST"},
	}
	RouterRule["canRoleListEdit"] = []UrlInfo{
		{Url: "/api/v1/role/:id", Method: "GET"},
		{Url: "/api/v1/role/:id", Method: "PUT"},
	}
	RouterRule["canRoleListDelete"] = []UrlInfo{
		{Url: "/api/v1/role", Method: "DELETE"},
	}

	// canDeptListGetPage
	RouterRule["canDeptListGetPage"] = []UrlInfo{
		{Url: "/api/v1/dept", Method: "GET"},
	}
	RouterRule["canDeptListNew"] = []UrlInfo{
		{Url: "/api/v1/dept", Method: "POST"},
	}
	RouterRule["canDeptListEdit"] = []UrlInfo{
		{Url: "/api/v1/dept/:id", Method: "GET"},
		{Url: "/api/v1/dept/:id", Method: "PUT"},
	}
	RouterRule["canDeptListDelete"] = []UrlInfo{
		{Url: "/api/v1/dept", Method: "DELETE"},
	}

	// canPostListGetPage
	RouterRule["canPostListGetPage"] = []UrlInfo{
		{Url: "/api/v1/post", Method: "GET"},
	}
	RouterRule["canPostListNew"] = []UrlInfo{
		{Url: "/api/v1/post", Method: "POST"},
	}
	RouterRule["canPostListEdit"] = []UrlInfo{
		{Url: "/api/v1/post/:id", Method: "GET"},
		{Url: "/api/v1/post/:id", Method: "PUT"},
	}
	RouterRule["canPostListDelete"] = []UrlInfo{
		{Url: "/api/v1/post", Method: "DELETE"},
	}

	// Process
	RouterRule["canProcessListGetPage"] = []UrlInfo{
		{Url: "/api/v1/process", Method: "GET"},
	}
	RouterRule["canProcessListNew"] = []UrlInfo{
		{Url: "/api/v1/process", Method: "POST"},
	}
	RouterRule["canProcessListEdit"] = []UrlInfo{
		{Url: "/api/v1/process/:id", Method: "GET"},
		{Url: "/api/v1/process/:id", Method: "PUT"},
	}
	RouterRule["canProcessListDelete"] = []UrlInfo{
		{Url: "/api/v1/process", Method: "DELETE"},
	}
	RouterRule["canProcessListClone"] = []UrlInfo{
		{Url: "/api/v1/process/clone/:id", Method: "POST"},
	}

	// WorkOrder
	RouterRule["canWorkOrderListGetPage"] = []UrlInfo{
		{Url: "/api/v1/work-order", Method: "GET"},
	}
	RouterRule["canWorkOrderListNew"] = []UrlInfo{
		{Url: "/api/v1/work-order", Method: "POST"},
	}
	RouterRule["canWorkOrderListEdit"] = []UrlInfo{
		{Url: "/api/v1/work-order/:id", Method: "GET"},
		{Url: "/api/v1/work-order/:id", Method: "PUT"},
	}
	RouterRule["canWorkOrderListDelete"] = []UrlInfo{
		{Url: "/api/v1/work-order", Method: "DELETE"},
	}
	RouterRule["canWorkListGet"] = []UrlInfo{
		{Url: "/api/v1/work", Method: "GET"},
	}
	RouterRule["canWorkListPost"] = []UrlInfo{
		{Url: "/api/v1/work", Method: "POST"},
	}
	RouterRule["canWorkListGet"] = []UrlInfo{
		{Url: "/api/v1/work/list", Method: "GET"},
	}

}

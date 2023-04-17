package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/graph/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerNodeRouter)
}

// registerNodeRouter
func registerNodeRouter(v1 *gin.RouterGroup) {
	labelApi := apis.LabelApi{}
	rLabel := v1.Group("/graph/tree")
	{
		rLabel.GET("/node", labelApi.GetNodeById)                                                     // get node by id
		rLabel.GET("/node/children", labelApi.GetChildrenNode)                                        // get children nodes by id
		rLabel.GET("/node/root", labelApi.GetLabelRootNode)                                           // get root node
		rLabel.GET("/path/between", labelApi.GetPathBetween)                                          // get TreeNodes(with Children) between two nodes
		rLabel.GET("/node/company/title/autoComplete", labelApi.GetCompanyTitleAutoCompleteByKeyWord) // get company title auto complete by keyword
		rLabel.GET("/path/company/title/fuzzyMatch", labelApi.FuzzyMatchCompanyFromSourceByTitle)     // get company fuzzy matched treeNode by title
		rLabel.GET("/node/label/title/autoComplete", labelApi.GetLabelTitleAutoCompleteByKeyWord)     // get label title auto complete by keyword
		rLabel.GET("/path/label/title/fuzzyMatch", labelApi.FuzzyMatchTagsFromSourceByTitle)          // get label fuzzy matched treeNode by title
	}

	linkApi := apis.LinkApi{}
	rLink := v1.Group("/graph/link")
	{
		rLink.GET("/node/root", linkApi.GetRootNode)          // get root node
		rLink.GET("/node", linkApi.GetNodeById)               // get node by id
		rLink.GET("/net/children", linkApi.GetNetToChildren)  // return children in struct Net of a node by given id
		rLink.GET("/net/parents", linkApi.GetNetToParents)    // return children in struct Net of a node by given id
		rLink.GET("/net/expand", linkApi.ExpandNetFromSource) // expand net from source by given id constraint by depth
	}
}

//			老版本的路由配置	--> 	新版本的路由配置
// /dev-api/graphDev/labels/root --> /dev-api/v1/graph/tree/node/root				无参数
// /dev-api/graphDev/labels/children --> /dev-api/v1/graph/tree/node/children		请求参数: id
// /dev-api/graphDev/labels/path --> /dev-api/v1/graph/tree/path/between			请求参数: sourceId(默认为rootId), targetId
// /dev-api/graphDev/labels/company/title/autoComplete --> /dev-api/v1/graph/tree/node/label/title/autoComplete	请求参数: pageSize, pageNum(从1开始), keyWord(原title)
// /dev-api/graphDev/labels/label/title/autoComplete --> /dev-api/v1/graph/tree/node/label/title/autoComplete		请求参数: pageSize, pageNum(从1开始), keyWord(原title)
// /dev-api/graphDev/labels/company/title/fuzzyMatch --> /dev-api/v1/graph/tree/path/company/title/fuzzyMatch		请求参数: pageSize, pageNum(从1开始), keyWord(原title)
// /dev-api/graphDev/labels/label/title/fuzzyMatch --> /dev-api/v1/graph/tree/path/label/title/fuzzyMatch			请求参数: pageSize, pageNum(从1开始), keyWord(原title)
//
// /dev-api/graphDev/link/node --> /dev-api/v1/graph/link/node				请求参数: id
// /dev-api/graphDev/link/root --> /dev-api/v1/graph/link/node/root		无参数
// /dev-api/graphDev/link/children --> /dev-api/v1/graph/link/net/children		请求参数: id, pageSize, pageNum(从1开始)
// /dev-api/graphDev/link/children --> /dev-api/v1/graph/link/net/parents		请求参数: id, pageSize, pageNum(从1开始)
// /dev-api/graphDev/link/path --> /dev-api/v1/graph/link/expand			请求参数: sourceId(默认为rootId, 原node_id), depth(默认为5), limit(默认为50)

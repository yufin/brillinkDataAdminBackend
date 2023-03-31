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
		rLabel.GET("/path/label/title/fuzzyMatch", labelApi.FuzzyMatchLabelsFromSourceByTitle)        // get label fuzzy matched treeNode by title
	}

	linkApi := apis.LinkApi{}
	rLink := v1.Group("/graph/link")
	{
		rLink.GET("/node/root", linkApi.GetRootNode)      // get root node
		rLink.GET("/node", linkApi.GetNodeById)           // get node by id
		rLink.GET("/expand", linkApi.ExpandNetFromSource) // expand net from source by given id constraint by depth
		rLink.GET("/children", linkApi.GetNetToChildren)  // return children in struct Net of a node by given id
	}
}

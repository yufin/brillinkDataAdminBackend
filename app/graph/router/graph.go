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
		rLabel.GET("/node", labelApi.GetNodeById)
		rLabel.GET("/node/children", labelApi.GetChildrenNode)
		rLabel.GET("/node/root", labelApi.GetLabelRootNode)

		rLabel.GET("/path/between", labelApi.GetPathBetween)

		rLabel.GET("/node/company/title/autoComplete", labelApi.GetCompanyTitleAutoCompleteByKeyWord)
		rLabel.GET("/node/label/title/autoComplete", labelApi.GetLabelTitleAutoCompleteByKeyWord)
	}

	linkApi := apis.LinkApi{}
	rLink := v1.Group("/graph/link")
	{
		rLink.GET("/node/root", linkApi.GetRootNode)
		rLink.GET("/expand", linkApi.ExpandNetFromSource)
		rLink.GET("/children", linkApi.GetNetToChildren)

	}

}

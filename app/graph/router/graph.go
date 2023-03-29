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
	rLabel := v1.Group("/graph/label")
	{
		rLabel.GET("/node", labelApi.GetNodeById)
		rLabel.GET("/node/children", labelApi.GetChildrenNode)
		rLabel.GET("/node/root", labelApi.GetLabelRootNode)
	}

	linkApi := apis.LinkApi{}
	rLink := v1.Group("/graph/link")
	{
		rLink.GET("/expand", linkApi.ExpandNetFromSource)
	}

}

package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/graph/constant"
	"go-admin/app/graph/service"
	"go-admin/app/graph/service/dto"
	"go-admin/common/apis"
	"net/http"
)

type LabelApi struct {
	apis.Api
}

func (e LabelApi) GetNodeById(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("id")
	nodeArr := service.GetNodeById(c.Request.Context(), id)
	if len(nodeArr) == 1 {
		e.OK(dto.SerializeTreeNode(nodeArr[0]))
	} else if len(nodeArr) > 1 {
		e.Error(http.StatusConflict, "More than one node found", "2")
	} else {
		var null *int = nil
		e.OK(null)
	}
}

func (e LabelApi) GetChildrenNode(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("id")
	children := service.GetChildrenById(c.Request.Context(), id, constant.LabelExpectRels)
	resp := make([]dto.TreeNode, 0)
	for _, child := range children {
		resp = append(resp, dto.SerializeTreeNode(child))
	}
	e.OK(resp)
}

func (e LabelApi) GetLabelRootNode(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	nodeArr := service.GetNodeById(c.Request.Context(), constant.LabelRootId)
	e.OK(dto.SerializeTreeNode(nodeArr[0]))
}

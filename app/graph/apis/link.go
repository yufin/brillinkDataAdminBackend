package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/graph/constant"
	"go-admin/app/graph/service"
	"go-admin/app/graph/service/dto"
	"go-admin/common/apis"
	"math"
	"strconv"
)

type LinkApi struct {
	apis.Api
}

func (e LinkApi) ExpandNetFromSource(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	var (
		depth   int
		limit   int
		errConv error
	)
	depth, errConv = strconv.Atoi(c.DefaultQuery("depth", "5"))
	if errConv != nil {
		depth = 5
	}
	limit, errConv = strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		limit = 50
	}
	result := service.ExpandPathFromSource(
		c.Request.Context(), c.DefaultQuery("sourceId", constant.LinkRootId), depth, limit)
	if len(result) == 0 {
		var null *int = nil
		e.OK(null)
	} else {
		resp := dto.SerializeNetFromPath(&result)
		e.OK(resp)
	}
}

func (e LinkApi) GetRootNode(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	nodeArr := service.GetNodeById(c.Request.Context(), constant.LinkRootId)
	if len(nodeArr) == 0 {
		var null *int = nil
		e.OK(null)
		return
	}
	e.OK(dto.SerializeLinkNode(nodeArr[0]))
}

func (e LinkApi) GetNetToChildren(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("sourceId")
	pageSize := c.DefaultQuery("pageSize", "50")
	pageNum := c.DefaultQuery("pageNum", "1")
	pageSizeInt, errConv := strconv.Atoi(pageSize)
	if errConv != nil {
		pageSizeInt = 10
	}
	pageNumInt, errConv := strconv.Atoi(pageNum)
	if errConv != nil {
		pageNumInt = 1
	}
	neoPath, total := service.GetPathToChildren(c.Request.Context(), id, pageSizeInt, pageNumInt)
	if len(neoPath) == 0 {
		var null *int = nil
		e.OK(null)
		return
	}
	resp := dto.SerializeNetFromPath(&neoPath)
	e.OK(dto.PaginatedResp{
		TotalPage: int(math.Ceil(float64(total) / float64(pageSizeInt))),
		PageNum:   pageNumInt,
		PageSize:  len(resp.Edges),
		Data:      resp,
	})
}

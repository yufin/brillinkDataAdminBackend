package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/graph/constant"
	"go-admin/app/graph/service"
	"go-admin/app/graph/service/dto"
	"go-admin/common/apis"
	"math"
	"net/http"
	"strconv"
)

type LinkApi struct {
	apis.Api
}

// ExpandNetFromSource 从某个节点开始，展开网络
// @Tags 图谱/网状图
// @Summary 通过给定节点和给定延伸层数检索延伸出的关联子节点，
// @Produce json
// @Param depth query string true "Depth of expansion"
// @Param limit query string true "Limit of each depth"
// @Param sourceId query string true "Source node id"
// @Success 200 {object} dto.Net
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/link/net/expand [get]
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
		depth = 3
	}
	depth = int(math.Min(float64(depth), 3))

	limit, errConv = strconv.Atoi(c.DefaultQuery("limit", "50"))
	if err != nil {
		limit = 50
	}

	s := service.PathService{}
	result, err := s.ExpandPathFromSource(
		c.Request.Context(), c.DefaultQuery("sourceId", constant.LinkRootId), depth, limit)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	if len(result) == 0 {
		var null *int = nil
		e.OK(null)
	} else {
		resp := dto.SerializeNetFromPath(&result)
		e.OK(resp)
	}
}

// GetRootNode 获取网状图的根节点
// @Tags 图谱/网状图
// @Summary 获取网状图的根节点
// @Produce json
// @Success 200 {object} dto.LinkNode
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/link/node/root [get]
func (e LinkApi) GetRootNode(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	s := service.NodeService{}
	nodeArr, err := s.GetNodeById(c.Request.Context(), constant.LinkRootId)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	if len(nodeArr) == 0 {
		var null *int = nil
		e.OK(null)
		return
	}
	e.OK(dto.SerializeLinkNode(nodeArr[0]))
}

// GetNodeById 通过id获取节点
// @Tags 图谱/网状图
// @Summary 通过id检索节点
// @Produce json
// @Param id query string true "Node id"
// @Success 200 {object} dto.LinkNode
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/link/node [get]
func (e LinkApi) GetNodeById(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("id")
	s := service.NodeService{}
	nodeArr, err := s.GetNodeById(c.Request.Context(), id)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	if len(nodeArr) == 0 {
		var null *int = nil
		e.OK(null)
		return
	}
	e.OK(dto.SerializeLinkNode(nodeArr[0]))
}

// GetNetToChildren 通过id获取节点的子节点
// @Tags 图谱/网状图
// @Summary 通过id检索节点的子节点
// @Produce json
// @Param id query string true "Node id"
// @Param pageSize query string true "Page size"
// @Param pageNum query string true "Page number"
// @Success 200 {object} dto.Net
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/link/net/children [get]
func (e LinkApi) GetNetToChildren(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("id")
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
	s := service.PathService{}
	neoPath, total, err := s.GetPathToChildren(c.Request.Context(), id, pageSizeInt, pageNumInt)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	if len(neoPath) == 0 {
		var null *int = nil
		e.OK(null)
		return
	}
	resp := dto.SerializeNetFromPath(&neoPath)
	e.PageOK(resp, int64(math.Ceil(float64(total)/float64(pageSizeInt))), pageNumInt, pageSizeInt)
}

// GetNetToParents 通过id检索节点的父节点
// @Tags 图谱/网状图
// @Summary 通过id检索节点的父节点
// @Produce json
// @Param id query string true "Node id"
// @Param pageSize query string true "Page size"
// @Param pageNum query string true "Page number"
// @Success 200 {object} dto.Net
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/link/net/parents [get]
func (e LinkApi) GetNetToParents(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("id")
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
	s := service.PathService{}
	neoPath, total, err := s.GetPathToParents(c.Request.Context(), id, pageSizeInt, pageNumInt)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	if len(neoPath) == 0 {
		var null *int = nil
		e.OK(null)
		return
	}
	resp := dto.SerializeNetFromPath(&neoPath)
	e.PageOK(resp, int64(math.Ceil(float64(total)/float64(pageSizeInt))), pageNumInt, pageSizeInt)
}

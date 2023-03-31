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
		resp = append(resp, *dto.SerializeTreeNode(child))
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

func (e LabelApi) GetCompanyTitleAutoCompleteByKeyWord(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	var (
		pageSize int
		pageNum  int
		errConv  error
	)
	pageSize, errConv = strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if errConv != nil {
		pageSize = 20
	}
	pageNum, errConv = strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		pageNum = 1
	}
	pageNum = int(math.Max(1, float64(pageNum)))
	pageSize = int(math.Max(1, float64(pageSize)))
	title := c.Query("title")
	cRes := make(chan []any)
	cTotal := make(chan int64)
	go func() {
		total := service.CountCompanyTitleAutoComplete(c.Request.Context(), title)
		cTotal <- total
	}()
	go func() {
		res := service.GetCompanyTitleAutoComplete(c.Request.Context(), title, pageSize, pageNum)
		cRes <- res
	}()
	titleList := <-cRes
	e.PageOK(titleList, int64(math.Ceil(float64(<-cTotal)/float64(pageSize))), pageNum, len(titleList))
}

func (e LabelApi) GetLabelTitleAutoCompleteByKeyWord(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	var (
		pageSize int
		pageNum  int
		errConv  error
	)
	pageSize, errConv = strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if errConv != nil {
		pageSize = 20
	}
	pageNum, errConv = strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		pageNum = 1
	}
	pageNum = int(math.Max(1, float64(pageNum)))
	pageSize = int(math.Max(1, float64(pageSize)))
	title := c.Query("title")
	cRes := make(chan []any)
	cTotal := make(chan int64)
	go func() {
		total := service.CountLabelTitleAutoComplete(c.Request.Context(), title)
		cTotal <- total
	}()
	go func() {
		res := service.GetLabelTitleAutoComplete(c.Request.Context(), title, pageSize, pageNum)
		cRes <- res
	}()
	titleList := <-cRes
	e.PageOK(titleList, int64(math.Ceil(float64(<-cTotal)/float64(pageSize))), pageNum, len(titleList))
}

// GetPathBetween 获取两个节点之间的TreeNode, sourceId默认为LabelRootId
func (e LabelApi) GetPathBetween(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	targetId := c.Query("targetId")
	sourceId := c.DefaultQuery("sourceId", constant.LabelRootId)
	filterStmt := "WHERE all(rel in relationships(p) WHERE not type(rel) in ['TRADING'])"
	neoPath := service.GetPathBetween(c.Request.Context(), sourceId, targetId, filterStmt)
	if len(neoPath) != 0 {
		resp := dto.SerializeTreeFromPath(&neoPath)
		e.OK(resp)
	}
}

func (e LabelApi) FuzzyMatchLabelsFromSourceByTitle(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	var (
		pageSize int
		pageNum  int
		errConv  error
	)
	pageSize, errConv = strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if errConv != nil {
		pageSize = 20
	}
	pageNum, errConv = strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		pageNum = 1
	}
	title := c.Query("title")

	cMatched := make(chan []any)
	cTotal := make(chan int64)
	go func() {
		rMatched := service.GetLabelTitleAutoComplete(c.Request.Context(), title, pageSize, pageNum)
		cMatched <- rMatched
	}()
	go func() {
		rTotal := service.CountLabelTitleAutoComplete(c.Request.Context(), title)
		cTotal <- rTotal
	}()

	matched := <-cMatched
	total := <-cTotal
	if len(matched) == 0 {
		e.PageOK(nil, total, pageNum, 0)
		return
	}
	ids := make([]string, 0)
	for _, item := range matched {
		ids = append(ids, item.(map[string]any)["id"].(string))
	}
	resp := service.GetPathFromSourceByIds(c.Request.Context(), constant.LabelRootId, ids, []string{"Label"}, constant.LabelExpectRels)
	e.PageOK(dto.SerializeTreeFromPath(&resp), int64(math.Ceil(float64(total)/float64(pageSize))), pageNum, len(matched))
}

func (e LabelApi) FuzzyMatchCompanyFromSourceByTitle(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	var (
		pageSize int
		pageNum  int
		errConv  error
	)
	pageSize, errConv = strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if errConv != nil {
		pageSize = 20
	}
	pageNum, errConv = strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		pageNum = 1
	}
	title := c.Query("title")

	cMatched := make(chan []any)
	cTotal := make(chan int64)
	go func() {
		rMatched := service.GetCompanyTitleAutoComplete(c.Request.Context(), title, pageSize, pageNum)
		cMatched <- rMatched
	}()
	go func() {
		rTotal := service.CountCompanyTitleAutoComplete(c.Request.Context(), title)
		cTotal <- rTotal
	}()

	matched := <-cMatched
	total := <-cTotal
	if len(matched) == 0 {
		e.PageOK(nil, total, pageNum, 0)
		return
	}
	ids := make([]string, 0)
	for _, item := range matched {
		ids = append(ids, item.(map[string]any)["id"].(string))
	}
	resp := service.GetPathFromSourceByIds(c.Request.Context(), constant.LabelRootId, ids, []string{"Company"}, constant.LabelExpectRels)
	e.PageOK(dto.SerializeTreeFromPath(&resp), int64(math.Ceil(float64(total)/float64(pageSize))), pageNum, len(matched))
}

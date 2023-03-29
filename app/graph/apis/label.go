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
	paginatedResp := dto.PaginatedResp{
		TotalPage: int(math.Ceil(float64(<-cTotal) / float64(pageSize))),
		PageNum:   pageNum,
		PageSize:  len(titleList),
		Data:      titleList,
	}
	e.OK(paginatedResp)
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
	paginatedResp := dto.PaginatedResp{
		TotalPage: int(math.Ceil(float64(<-cTotal) / float64(pageSize))),
		PageNum:   pageNum,
		PageSize:  len(titleList),
		Data:      titleList,
	}
	e.OK(paginatedResp)
}

func (e LabelApi) GetPathToRoot(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("nodeId")
	filterStmt := "WHERE all(rel in relationships(p) WHERE not type(rel) in ['TRADING'])"
	path := service.GetPathBetween(c.Request.Context(), constant.LabelRootId, id, filterStmt)
	e.OK(path)
}

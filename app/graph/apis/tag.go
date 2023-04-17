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

// GetNodeById handles GET requests to retrieve a single node by its ID.
// This method is used by go-swagger.
// @Tags 图谱/树状图
// @Summary 通过id检索单个节点
// @Produce json
// @Param id query string true "Node ID"
// @Success 200 {object} dto.TreeNode
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/tree/node [get]
func (e LabelApi) GetNodeById(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("id")
	nodeArr, err := service.GetNodeById(c.Request.Context(), id)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	if len(nodeArr) == 0 {
		var null *int = nil
		e.OK(null)
		return
	}
	e.OK(dto.SerializeTreeNode(nodeArr[0]))
}

// GetLabelRootNode handles GET requests to retrieve the root node of the label tree.
// This method is used by go-swagger.
// @Tags 图谱/树状图
// @Summary 查询标签树的根节点
// @Produce json
// @Success 200 {object} dto.TreeNode
// @Failure 204 {object} string
// @Router /dev-api/v1/graph/tree/node/root [get]
func (e LabelApi) GetLabelRootNode(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	nodeArr, err := service.GetNodeById(c.Request.Context(), constant.LabelRootId)
	if err != nil {
		e.Error(http.StatusNoContent, err.Error(), "1")
		return
	}
	if len(nodeArr) == 0 {
		var null *int = nil
		e.OK(null)
		return
	}
	e.OK(dto.SerializeTreeNode(nodeArr[0]))
}

// GetChildrenNode handles GET requests to retrieve the children nodes of a given node.
// This method is used by go-swagger.
// @Tags 图谱/树状图
// @Summary 检索给定节点的子节点
// @Produce json
// @Param id query string true "Node ID"
// @Param pageSize query int false "Page size"
// @Param pageNum query int false "Page number"
// @Success 200 {array} dto.TreeNode
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/tree/node/children [get]
func (e LabelApi) GetChildrenNode(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	id := c.Query("id")
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
	children, count, err := service.GetChildrenById(c.Request.Context(), id, pageSize, pageNum, constant.LabelExpectRels)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	var resp []dto.TreeNode
	for _, child := range children {
		childNode := dto.SerializeTreeNode(child)
		resp = append(resp, *childNode)
	}
	e.PageOK(resp, int64(math.Ceil(float64(count)/float64(pageSize))), pageNum, len(resp))
}

// GetCompanyTitleAutoCompleteByKeyWord handles GET requests to retrieve a list of company titles that match a given keyword for autocomplete purposes.
// This method is used by go-swagger.
// @Tags 图谱/树状图
// @Summary 通过给定的关键词检索模糊匹配的公司名称列表
// @Produce json
// @Param keyWord query string true "Keyword to search for"
// @Param pageSize query int false "Page size"
// @Param pageNum query int false "Page number"
// @Success 200 {array} any
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/tree/node/company/title/autoComplete [get]
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
	title := c.Query("keyWord")
	cRes := make(chan struct {
		result []any
		err    error
	})
	cTotal := make(chan struct {
		result int64
		err    error
	})
	go func() {
		rTotal, err := service.CountCompanyTitleAutoComplete(c.Request.Context(), title)
		cTotal <- struct {
			result int64
			err    error
		}{result: rTotal, err: err}
	}()
	go func() {
		res, err := service.GetCompanyTitleAutoComplete(c.Request.Context(), title, pageSize, pageNum)
		cRes <- struct {
			result []any
			err    error
		}{result: res, err: err}
	}()
	titleList := <-cRes
	total := <-cTotal
	if total.err != nil {
		e.Error(http.StatusInternalServerError, total.err.Error(), "1")
		return
	}
	if titleList.err != nil {
		e.Error(http.StatusInternalServerError, titleList.err.Error(), "1")
		return
	}
	e.PageOK(titleList.result, int64(math.Ceil(float64(total.result)/float64(pageSize))), pageNum, len(titleList.result))
}

// GetLabelTitleAutoCompleteByKeyWord handles GET requests to retrieve a list of label titles that match a given keyword for autocomplete purposes.
// This method is used by go-swagger.
// @Tags 图谱/树状图
// @Summary 通过给定的关键词检索模糊匹配的标签名称列表
// @Produce json
// @Param keyWord query string true "Keyword to search for"
// @Param pageSize query int false "Page size"
// @Param pageNum query int false "Page number"
// @Success 200 {array} any
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/tree/node/label/title/autoComplete [get]
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
	title := c.Query("keyWord")
	cRes := make(chan struct {
		result []any
		err    error
	})
	cTotal := make(chan struct {
		result int64
		err    error
	})
	go func() {
		rTotal, err := service.CountLabelTitleAutoComplete(c.Request.Context(), title)
		cTotal <- struct {
			result int64
			err    error
		}{result: rTotal, err: err}
	}()
	go func() {
		res, err := service.GetLabelTitleAutoComplete(c.Request.Context(), title, pageSize, pageNum)
		cRes <- struct {
			result []any
			err    error
		}{result: res, err: err}
	}()
	titleList := <-cRes
	total := <-cTotal
	if total.err != nil {
		e.Error(http.StatusInternalServerError, total.err.Error(), "1")
		return
	}
	if titleList.err != nil {
		e.Error(http.StatusInternalServerError, titleList.err.Error(), "1")
		return
	}
	e.PageOK(titleList.result, int64(math.Ceil(float64(total.result)/float64(pageSize))), pageNum, len(titleList.result))
}

// GetPathBetween handles GET requests to retrieve the path between two labels.
// @Tags 图谱/树状图
// @Summary 获取标签数图中两个节点之间的路径
// @Produce json
// @Param sourceId query string false "Source label ID, default to root label ID"
// @Param targetId query string true "Target label ID"
// @Success 200 {object} any
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/tree/path/between [get]
func (e LabelApi) GetPathBetween(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	targetId := c.Query("targetId")
	sourceId := c.DefaultQuery("sourceId", constant.LabelRootId)
	filterStmt := "WHERE all(rel in relationships(p) WHERE not type(rel) in ['TRADING'])"
	neoPath, err := service.GetPathBetween(c.Request.Context(), sourceId, targetId, filterStmt)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	if len(neoPath) != 0 {
		resp := dto.SerializeTreeFromPath(&neoPath)
		e.OK(resp)
	}
}

// FuzzyMatchTagsFromSourceByTitle fuzzy match Tags from source by title
// This annotation is used by go-swagger.
// @Tags 图谱/树状图
// @Summary 通过标签名称关键词模糊匹配标签节点
// @Produce json
// @Param sourceId query string false "Source label ID, default to root label ID"
// @Param targetId query string true "Target label ID"
// @Success 200 {object} any
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/tree/path/label/title/fuzzyMatch [get]
func (e LabelApi) FuzzyMatchTagsFromSourceByTitle(c *gin.Context) {
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
	title := c.Query("keyWord")

	cMatched := make(chan struct {
		result []any
		err    error
	})
	cTotal := make(chan struct {
		result int64
		err    error
	})
	go func() {
		rMatched, err := service.GetLabelTitleAutoComplete(c.Request.Context(), title, pageSize, pageNum)
		cMatched <- struct {
			result []any
			err    error
		}{result: rMatched, err: err}
	}()
	go func() {
		rTotal, err := service.CountLabelTitleAutoComplete(c.Request.Context(), title)
		cTotal <- struct {
			result int64
			err    error
		}{result: rTotal, err: err}
	}()

	matched := <-cMatched
	total := <-cTotal
	if total.err != nil {
		e.Error(http.StatusInternalServerError, total.err.Error(), "1")
		return
	}
	if matched.err != nil {
		e.Error(http.StatusInternalServerError, matched.err.Error(), "1")
		return
	}
	if len(matched.result) == 0 {
		e.PageOK(nil, total.result, pageNum, 0)
		return
	}
	ids := make([]string, 0)
	for _, item := range matched.result {
		ids = append(ids, item.(map[string]any)["id"].(string))
	}
	resp, err := service.GetPathFromSourceByIds(c.Request.Context(), constant.LabelRootId, ids, []string{"Label"}, constant.LabelExpectRels)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	e.PageOK(dto.SerializeTreeFromPath(&resp), int64(math.Ceil(float64(total.result)/float64(pageSize))), pageNum, len(matched.result))
}

// FuzzyMatchCompanyFromSourceByTitle fuzzy match company from source by title
// This method is used by go-swagger.
// @Tags 图谱/树状图
// @Summary 通过公司名称关键词模糊匹配公司节点
// @Produce json
// @Param sourceId query string false "Source label ID, default to root label ID"
// @Param targetId query string true "Target label ID"
// @Success 200 {object} any
// @Failure 500 {object} string
// @Router /dev-api/v1/graph/tree/path/company/title/fuzzyMatch [get]
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
	pageSize, errConv = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if errConv != nil {
		pageSize = 10
	}
	pageNum, errConv = strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		pageNum = 1
	}
	title := c.Query("keyWord")

	cMatched := make(chan struct {
		result []any
		err    error
	})
	cTotal := make(chan struct {
		result int64
		err    error
	})
	go func() {
		rMatched, err := service.GetCompanyTitleAutoComplete(c.Request.Context(), title, pageSize, pageNum)
		cMatched <- struct {
			result []any
			err    error
		}{result: rMatched, err: err}
	}()
	go func() {
		rTotal, err := service.CountCompanyTitleAutoComplete(c.Request.Context(), title)
		cTotal <- struct {
			result int64
			err    error
		}{result: rTotal, err: err}
	}()

	matched := <-cMatched
	if matched.err != nil {
		e.Error(http.StatusInternalServerError, matched.err.Error(), "1")
		return
	}
	total := <-cTotal
	if total.err != nil {
		e.Error(http.StatusInternalServerError, total.err.Error(), "1")
		return
	}
	if len(matched.result) == 0 {
		e.PageOK(nil, total.result, pageNum, 0)
		return
	}
	ids := make([]string, 0)
	for _, item := range matched.result {
		ids = append(ids, item.(map[string]any)["id"].(string))
	}
	resp, err := service.GetPathFromSourceByIds(c.Request.Context(), constant.LabelRootId, ids, []string{"Company"}, constant.LabelExpectRels)
	if err != nil {
		e.Error(http.StatusInternalServerError, err.Error(), "1")
		return
	}
	e.PageOK(dto.SerializeTreeFromPath(&resp), int64(math.Ceil(float64(total.result)/float64(pageSize))), pageNum, len(matched.result))
}

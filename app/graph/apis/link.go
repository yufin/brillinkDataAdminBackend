package apis

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/graph/service"
	"go-admin/app/graph/service/dto"
	"go-admin/common/apis"
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
	result := service.GetPathFromSource(c.Request.Context(), c.Query("sourceId"), depth, limit)
	if len(result) == 0 {
		var null *int = nil
		e.OK(null)
	} else {
		resp := dto.SerializeNetFromPath(&result)
		e.OK(resp)
	}

}

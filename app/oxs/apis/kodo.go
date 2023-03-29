/**
* @Author: Akiraka
* @Date: 2022/10/17 11:41
 */

package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/oxs/models"
	"go-admin/app/oxs/service"
)

// GetKodo 七牛云对象存储
func (e OXS) GetKodo(c *gin.Context) {
	e.MakeContext(c)
	s := service.OXS{}
	res := s.GetKodo(c)
	e.OK(models.ResponseKODO{
		Enable:       true,
		OxsType:      sdk.Runtime.GetConfig(c.Request.Host, "oxs_type").(string),
		Region:       sdk.Runtime.GetConfig(c.Request.Host, "oxs_region").(string),
		AccessDomain: sdk.Runtime.GetConfig(c.Request.Host, "oxs_access_domain").(string),
		Bucket:       sdk.Runtime.GetConfig(c.Request.Host, "oxs_bucket").(string),
		Token:        res,
		Status:       true,
	})
}

package server

import (
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/jobs/router"
)

func init() {
	sdk.Runtime.SetAppRouters(router.InitRouter)
}

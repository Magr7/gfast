package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonService "github.com/tiger1103/gfast/v3/internal/app/common/service"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/oncall", func(group *ghttp.RouterGroup) {
		group.Middleware(commonService.Middleware().MiddlewareCORS)
	})
}

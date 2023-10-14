package router

import (
	"github.com/gin-gonic/gin"
	"leetroll/common/runtime"
	"log"
	"os"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(*gin.RouterGroup), 0)
)

func InitAppRouter() {
	var r *gin.Engine
	h := runtime.App.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support regular engine")
		os.Exit(-1)
	}

	//todo 加入鉴权中间件

	//注册业务路由

	//1. 不需要认证的路由
	noCheckRoleRouter(r)

	//2. todo: 需要认证的路由
}

func noCheckRoleRouter(r *gin.Engine) {

	v1 := r.Group("app")

	for _, f := range routerNoCheckRole {
		f(v1)
	}

	for _, f := range routerCheckRole {
		f(v1)
	}
}

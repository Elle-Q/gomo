package router

import (
	"github.com/gin-gonic/gin"
	"gomo/common/runtime"
	"log"
	"os"
)

var routerCheckRole = make([]func(*gin.RouterGroup), 0)


func InitAdminRouter()  {
	var r *gin.Engine
	h := runtime.App.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine :
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support regular engine")
		os.Exit(-1)
	}

	//注册管理员路由
	checkRoleRouter(r)
}

func checkRoleRouter(r *gin.Engine) {

	v1 := r.Group("admin")
	for _,f := range routerCheckRole{
		f(v1)
	}
}
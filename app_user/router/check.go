package router

import (
	"github.com/gin-gonic/gin"
)

func init()  {
	routerNoCheckRole = append(routerNoCheckRole, registerRouter)

}

func registerRouter(g *gin.RouterGroup) {
	//
	//r := g.Group("/sysjob").Use(actions.PermissionAction())
	//{
	//	//r.GET("", api.GetPage)
	//	//r.GET("/:id", api.Get)
	//	//r.POST("", api.Insert)
	//	//r.PUT("", api.Update)
	//	//r.DELETE("", api.Delete)
	//}

}
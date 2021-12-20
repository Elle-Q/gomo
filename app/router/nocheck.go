package router

import (
	"github.com/gin-gonic/gin"
	"gomo/app/api"
	"gomo/common/actions"
)


func init()  {
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckRouter)
}

func registerNoCheckRouter(g *gin.RouterGroup) {

	_CatApi := api.Category{}
	category := g.Group("/category").Use(actions.PermissionAction())
	{
		category.GET("/list", _CatApi.List)
		category.GET("/:id", _CatApi.Get)
		//r.POST("", api.Insert)
		//r.PUT("", api.Update)
		//r.DELETE("", api.Delete)
	}

	_UserApi := api.User{}
	user := g.Group("/user").Use(actions.PermissionAction())
	{
		user.GET("/login", _UserApi.Login)
		user.GET("/refresh", _UserApi.Refresh)
	}
}
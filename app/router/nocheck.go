package router

import (
	"github.com/gin-gonic/gin"
	"leetroll/app/api"
	"leetroll/common/actions"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerNoCheckRouter)
}

func registerNoCheckRouter(g *gin.RouterGroup) {

	_CatApi := api.Category{}
	category := g.Group("/category").Use(actions.PermissionAction())
	{
		category.GET("/list", _CatApi.List)
		category.GET("/:id", _CatApi.Get)
		category.GET("/list/items", _CatApi.ListCatsWithItems)
		//r.POST("", api.Insert)
		//r.PUT("", api.Update)
		//r.DELETE("", api.Delete)
	}

	_UserApi := api.User{}
	user := g.Group("/user").Use(actions.PermissionAction())
	{
		user.POST("/login", _UserApi.Login)
		user.POST("/refresh", _UserApi.Refresh)
	}

	_ConfigApi := api.Config{}
	config := g.Group("/config").Use(actions.PermissionAction())
	{
		config.GET("/:name", _ConfigApi.FindDefaultAvatarByName)
	}

}

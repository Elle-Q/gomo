package router

import (
	"github.com/gin-gonic/gin"
	"gomo/admin/api"
	"gomo/common/middleware"
)

func init()  {
	routerCheckRole = append(routerCheckRole, registerCheckRouter)
}

func registerCheckRouter(g *gin.RouterGroup) {

	_UserApi := api.User{}
	user := g.Group("/user").Use(middleware.AuthJWTMiddleware())
	{
		user.GET("/:id", _UserApi.GetUser)
		user.GET("/list", _UserApi.List)
		//r.POST("", api.Insert)
		//r.PUT("", api.Update)
		//r.DELETE("", api.Delete)
	}

	_CatApi := api.Category{}
	cat := g.Group("/cat").Use(middleware.AuthJWTMiddleware())
	{
		//cat.GET("/:id", _UserApi.GetUser)
		cat.GET("/list", _CatApi.List)
		cat.GET("/list-name", _CatApi.ListName)
		cat.POST("/update", _CatApi.Update)
		cat.POST("/delete", _CatApi.Delete)
	}

	_ItemApi := api.Item{}
	item := g.Group("/item").Use(middleware.AuthJWTMiddleware())
	{
		item.GET("/list", _ItemApi.List)
	}

}
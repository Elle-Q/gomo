package router

import (
	"github.com/gin-gonic/gin"
	"gomo/admin/api"
	"gomo/common/actions"
)

func init()  {
	routerCheckRole = append(routerCheckRole, registerCheckRouter)
}

func registerCheckRouter(g *gin.RouterGroup) {

	_UserApi := api.User{}
	user := g.Group("/user").Use(actions.PermissionAction())
	{
		user.GET("/:id", _UserApi.GetUser)
		user.GET("/list", _UserApi.List)
		//r.POST("", api.Insert)
		//r.PUT("", api.Update)
		//r.DELETE("", api.Delete)
	}

	_CatApi := api.Category{}
	cat := g.Group("/cat").Use(actions.PermissionAction())
	{
		//cat.GET("/:id", _UserApi.GetUser)
		cat.GET("/list", _CatApi.List)
		cat.POST("/update", _CatApi.Save)
		cat.POST("/delete", _CatApi.Delete)
	}

}
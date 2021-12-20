package router

import (
	"github.com/gin-gonic/gin"
	"gomo/app/api"
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
		user.GET("/update", _UserApi.UpdateUser)
		user.GET("/logout", _UserApi.Logout)
		//r.GET("/:id", api.Get)
		//r.POST("", api.Insert)
		//r.PUT("", api.Update)
		//r.DELETE("", api.Delete)
	}


}
package router

import (
	"github.com/gin-gonic/gin"
	"gomo/app/api"
	"gomo/common/actions"
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
		user.POST("/update", _UserApi.UpdateUser)
		user.POST("/avatar/update", _UserApi.UpdateUserAvatar)
		user.POST("/bg/update", _UserApi.UpdateUserBG)
		user.GET("/logout", _UserApi.Logout)
	}

	_QiniuApi := api.Qiniu{}
	qiniu := g.Group("/qiniu").Use(actions.PermissionAction())
	{
		qiniu.GET("/token", _QiniuApi.GetUpToken)
	}


}
package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	// 数据库链接
	r.Use(WithContextDb)
}
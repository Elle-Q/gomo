package actions

import (
	"github.com/gin-gonic/gin"
)

func PermissionAction() gin.HandlerFunc {
	
	return func(c *gin.Context) {
		c.Next()
	}
	
}
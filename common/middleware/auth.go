package middleware

import (
	"github.com/gin-gonic/gin"
	"gomo/auth"
	"gomo/common/response"
	"net/http"
)

func  AuthJWTMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {

		err := auth.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		tokenAuth, err := auth.ExtractTokenMetadata(c.Request)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err,"unauthorized")
			return
		}
		_, err = auth.FetchAuth(tokenAuth)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err,"unauthorized")
			return
		}
		c.Next()
	}
}
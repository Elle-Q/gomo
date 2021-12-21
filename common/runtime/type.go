package runtime

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"net/http"
)

type Runtime interface {
	SetDb(db *sql.DB)
	GetDb() *sql.DB

	SetRedis(client *redis.Client)
	GetRedis() *redis.Client

	// SetEngine 使用的路由
	SetEngine(engine http.Handler)
	GetEngine() http.Handler

	GetRouter() []Router

	// SetMiddleware middleware
	SetMiddleware(string, interface{})
	GetMiddleware() map[string]interface{}
	GetMiddlewareKey(key string) interface{}

	SetHandler(key string, routerGroup func(r *gin.RouterGroup, hand ...*gin.HandlerFunc))
	GetHandler() map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)
	GetHandlerPrefix(key string) []func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)

}

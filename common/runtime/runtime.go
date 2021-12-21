package runtime

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"net/http"
	"sync"
)

type Application struct {
	db          *sql.DB
	engine      http.Handler
	mux         sync.RWMutex
	middlewares map[string]interface{}
	handler     map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)
	routers     []Router
	redis       *redis.Client
}

type Router struct {
	HttpMethod, RelativePath, Handler string
}

type Routers struct {
	List []Router
}

// SetDb 设置对应key的db
func (e *Application) SetDb(db *sql.DB) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.db = db
}

// GetDb 获取所有map里的db数据
func (e *Application) GetDb() *sql.DB {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.db
}

// SetRedis 设置redis
func (e *Application) SetRedis(redis *redis.Client) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.redis = redis
}

// GetRedis 获取redis
func (e *Application) GetRedis() *redis.Client {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.redis
}

// SetEngine 设置路由引擎
func (e *Application) SetEngine(engine http.Handler) {
	e.engine = engine
}

// GetEngine 获取路由引擎
func (e *Application) GetEngine() http.Handler {
	return e.engine
}

// GetRouter 获取路由表
func (e *Application) GetRouter() []Router {
	return e.setRouter()
}

// setRouter 设置路由表
func (e *Application) setRouter() []Router {
	switch e.engine.(type) {
	case *gin.Engine:
		routers := e.engine.(*gin.Engine).Routes()
		for _, router := range routers {
			e.routers = append(e.routers, Router{RelativePath: router.Path, Handler: router.Handler, HttpMethod: router.Method})
		}
	}
	return e.routers
}

// SetMiddleware 设置中间件
func (e *Application) SetMiddleware(key string, middleware interface{}) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.middlewares[key] = middleware
}

// GetMiddleware 获取所有中间件
func (e *Application) GetMiddleware() map[string]interface{} {
	return e.middlewares
}

// GetMiddlewareKey 获取对应key的中间件
func (e *Application) GetMiddlewareKey(key string) interface{} {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.middlewares[key]
}

func (e *Application) SetHandler(key string, routerGroup func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.handler[key] = append(e.handler[key], routerGroup)
}

func (e *Application) GetHandler() map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc) {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.handler
}

func (e *Application) GetHandlerPrefix(key string) []func(r *gin.RouterGroup, hand ...*gin.HandlerFunc) {
	e.mux.Lock()
	defer e.mux.Unlock()
	return e.handler[key]
}

func NewConfig() *Application {
	return &Application{
		middlewares: make(map[string]interface{}),
		handler:     make(map[string][]func(r *gin.RouterGroup, hand ...*gin.HandlerFunc)),
		routers:     make([]Router, 0),
	}
}

var App Runtime = NewConfig()

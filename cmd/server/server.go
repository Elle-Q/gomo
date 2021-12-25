package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	adminRouter "gomo/admin/router"
	appRouter "gomo/app/router"
	"gomo/common/global"
	"gomo/common/middleware"
	"gomo/common/runtime"
	"gomo/config"
	"gomo/tool"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	ServerStartCmd = &cobra.Command{
		Use:   "server",
		Short: "run server",
		Long:  `run server (user)`,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		Run: func(cmd *cobra.Command, args []string)  {
			run()
		},
	}
)

var AppRouters = make([]func(), 0)
var AdminRouters = make([]func(), 0)

func init() {

	//注册路由 fixme 其他应用的路由，在本目录新建文件放在init方法
	AppRouters = append(AppRouters, appRouter.InitAppRouter)

	AdminRouters = append(AdminRouters, adminRouter.InitAdminRouter)
}

//初始化操作
func setup() {

	//1. 读取配置, 初始化数据库
	config.Setup("config/config.yml")

}

//执行
func run()  {
	if config.ApplicationConfig.Mode == tool.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}

	runtime.App.SetEngine(gin.New())

	//1. 初始化数据库
	initDB()

	//3. 初始化redis
	initRedis()

	//3.路由, 中间件配置
	initRouters()

	for _, f := range AppRouters {
		f()
	}

	for _, f := range AdminRouters {
		f()
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler: runtime.App.GetEngine(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//新起线程
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("listen: ", err)
		}
	}()

	fmt.Println(tool.Red(string(global.Banner)))
	fmt.Println(tool.Green("Server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", config.ApplicationConfig.Port)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", tool.GetLocalHost(), config.ApplicationConfig.Port)

	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", tool.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", tool.GetCurrentTimeStr())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	//return nil
}

//初始化数据库
func initDB() {
	_cfg := config.DatabaseConfig

	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%v",
			_cfg.Host,
			_cfg.Port,
			_cfg.User,
			_cfg.Password,
			_cfg.Dbname,
			_cfg.SSLMode))

	if err != nil {
		fmt.Println(err.Error())
	}
	//defer db.Close()

	//设置全局数据库连接
	runtime.App.SetDb(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

//初始化redis
func initRedis() {
	_cfg := config.RedisConfig

	//Initializing redis

	client := redis.NewClient(&redis.Options{
		Addr: _cfg.Address,
		Password: _cfg.Password,
		DB: _cfg.DB,
	})

	//设置全局
	runtime.App.SetRedis(client)


	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

//初始化路由
func initRouters() {
	var r *gin.Engine
	h := runtime.App.GetEngine()
	if h == nil {
		h = gin.New()
		runtime.App.SetEngine(h)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	middleware.InitMiddleware(r)
}

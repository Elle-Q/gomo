package config

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v7"
	_ "github.com/lib/pq"
	"leetroll/common/runtime"
)

// 初始化数据库
func InitDB() {
	_cfg := DatabaseConfig

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

// 初始化redis
func InitRedis() {
	_cfg := RedisConfig

	//Initializing redis

	client := redis.NewClient(&redis.Options{
		Addr:     _cfg.Address,
		Password: _cfg.Password,
		DB:       _cfg.DB,
	})

	//设置全局
	runtime.App.SetRedis(client)

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

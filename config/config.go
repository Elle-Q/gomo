package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// 全部配置项
type Config struct {
	Application *Application
	Database    *Database
	JWT         *JWT
	Redis       *Redis
	Qiniu       *Qiniu
	//Logger      *Logger
	//Queue       *Queue
}

func Setup(filePath string) {
	_cfg := Config{
		ApplicationConfig,
		DatabaseConfig,
		JWTConfig,
		RedisConfig,
		QiniuConfig,
	}

	ymlData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(ymlData, &_cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}

package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)


//全部配置项
type Config struct {
	Application *Application
	Database    *Database
	JWT  *JWT
	Redis *Redis
	//Logger      *Logger
	//Queue       *Queue
}


func Setup(filePath string) {
	_cfg := Config{
		ApplicationConfig,
		DatabaseConfig,
		JWTConfig,
		RedisConfig,
	}

	ymlData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(ymlData, &_cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}

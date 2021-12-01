package config

import (
	"github.com/jinzhu/configor"
	"log"
)

// Settings 兼容原先的配置结构
type Settings struct {
	Settings  Config `yaml:"settings"`
	callbacks []func()
}

func (e *Settings) runCallback()  {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Settings) onChange() {
	e.init()
	log.Println("config change and reload")
}

func (e *Settings) init() {
	//e.Settings
}


type Config struct {
	Application *Application          `yaml:"application"`
	Logger      *Logger               `yaml:"logger"`
	Database    *Database             `yaml:"database"`
	//Cache       *Cache                `yaml:"cache"`
	//Queue       *Queue                `yaml:"queue"`
}

var Main = (func() Config  {

	var conf Config
	if err := configor.Load(&conf, "PATH_TO_CONFIG_FILE"); err != nil {
		panic(err.Error())
	}
	return conf
})()
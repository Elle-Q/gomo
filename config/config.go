package config

import (
	"github.com/jinzhu/configor"
	"gomo/common/file"
	"log"
)

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
	e.Settings.Logger.Setup()
	e.runCallback()
}


type Config struct {
	Application *Application          `yaml:"application"`
	Logger      *Logger               `yaml:"logger"`
	Database    *Database             `yaml:"database"`
	//Cache       *Cache                `yaml:"cache"`
	//Queue       *Queue                `yaml:"queue"`
}

func Setup(s file.Source, fs ...func())  {
	_cfg = &Settings{
		Settings: Config{
			Application: ApplicationConfig,
			Logger:      LoggerConfig,
		},
		callbacks: fs,
	}
	var err error

}

var Main = (func() Config  {

	var conf Config
	if err := configor.Load(&conf, "PATH_TO_CONFIG_FILE"); err != nil {
		panic(err.Error())
	}
	return conf
})()
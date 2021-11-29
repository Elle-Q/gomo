package config

import (
	"github.com/jinzhu/configor"
)

type Database struct {
	Host     string `default:"localhost"`
	Port     int `default:"5432"`
	User     string `default:"elle"`
	Password string `default:"Ggg123654."`
	Dbname   string `default:"homo"`
	SSLMode  bool `default:"true"`
}

type Configuration struct {
	db Database
}

var Main = (func() Configuration {

	var conf Configuration
	if err := configor.Load(&conf, "PATH_TO_CONFIG_FILE"); err != nil {
		panic(err.Error())
	}
	return conf
})()
package config

type Application struct {
	ReadTimeout   int    `yaml:"readtimeout"`
	WriterTimeout int    `yaml:"writertimeout"`
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Name          string `yaml:"name"`
	Mode          string `yaml:"mode"`
}

var ApplicationConfig = new(Application)

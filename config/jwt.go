package config



type JWT struct {
	AccessSecret   string    `yaml:"accessSecret"`
	RefreshSecret   string    `yaml:"refreshSecret"`
	Timeout int    `yaml:"timeout"`
}

var JWTConfig = new(JWT)

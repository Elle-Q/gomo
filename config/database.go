package config

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
	Driver   string `yaml:"driver"`
	SSLMode  string `yaml:"sslMode"`
	Dbname   string `yaml:"dbname"`
}

var (
	DatabaseConfig = new(Database)
)

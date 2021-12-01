package config

type Database struct {
	Host     string `default:"localhost"`
	Port     int `default:"5432"`
	User     string `default:"elle"`
	Password string `default:"Ggg123654."`
	Dbname   string `default:"homo"`
	SSLMode  bool `default:"true"`
}

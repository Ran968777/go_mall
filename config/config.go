package config

type Config struct {
	MySQL  MySQLConfig
	Server ServerConfig
}

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type ServerConfig struct {
	Port int
}

// GetConfig 返回应用配置
func GetConfig() *Config {
	return &Config{
		MySQL: MySQLConfig{
			Host:     "127.0.0.1",
			Port:     3306,
			User:     "root",
			Password: "123456",
			DBName:   "yami_shops",
		},
		Server: ServerConfig{
			Port: 8087,
		},
	}
}

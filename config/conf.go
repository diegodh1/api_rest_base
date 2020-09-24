package config

//Config struct
type Config struct {
	DB *DBConfig
}

//DBConfig struct
type DBConfig struct {
	Username string
	Password string
	Database string
	Port     int
	Host     string
}

//GetConfig returns a db configuration
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "riyaosewvtylgr",
			Password: "2fa70887881bf743c59b194a868bc1224c7f65b04fbc40c8512138e35e9a9d22",
			Database: "d64v91kmre708t",
			Port:     5432,
			Host:     "ec2-3-215-207-12.compute-1.amazonaws.com",
		},
	}
}

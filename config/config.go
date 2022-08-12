package config

import (
	"os"
)

// AppConfig main struct
type AppConfig struct {
	App struct {
		Host string
		Port string
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

var config *AppConfig

// Get untuk ngambil config
func Get() *AppConfig {

	if config == nil {
		config = &AppConfig{}
		config.App.Host = os.Getenv("APP_HOST")
		config.App.Port = os.Getenv("APP_PORT")
		config.Database.Host = os.Getenv("DB_HOST")
		config.Database.Port = os.Getenv("DB_PORT")
		config.Database.User = os.Getenv("DB_USER")
		config.Database.Pass = os.Getenv("DB_PASS")
		config.Database.Name = os.Getenv("DB_NAME")
	}
	return config
}
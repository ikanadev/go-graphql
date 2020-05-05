package config

import "os"

// this package just reads an maps the environment variables

// db Stores all db variables
type db struct {
	Username string
	Database string
	Password string
	Host     string
	Port     string
}

// gin all gin configuration variables
type gin struct {
	Mode string
}

// app all app config variables
type app struct {
	JWT  string
	Port string
}

// Config main struct
type Config struct {
	DB  db
	Gin gin
	App app
}

// GetConfig get a new config with all variables
func GetConfig() Config {
	return Config{
		DB: db{
			Username: getEnv("DB_USER", "postgres"),
			Database: getEnv("DB_DATABASE", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
		},
		Gin: gin{Mode: "debug"},
		App: app{
			JWT:  getEnv("JWT_KEY", ""),
			Port: getEnv("PORT", "8080"),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultVal
}

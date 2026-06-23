package config

import "os"

type Config struct {
	Port         string
	IsProduction bool
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	isProd := false
	if os.Getenv("ENV") == "production" {
		isProd = true
	}

	return Config{Port: port, IsProduction: isProd}
}

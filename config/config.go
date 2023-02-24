package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type (
	// Config -.
	Config struct {
		HTTP `json:"http"`
		Log  `json:"logger"`
		PG   `json:"postgres"`
	}

	// HTTP -.
	HTTP struct {
		HTTPPort string `json:"http_port"`
	}

	// Log -.
	Log struct {
		Level string `json:"level"`
	}

	// PG -.
	PG struct {
		Address  string `json:"address"`
		Port     int    `json:"port"`
		Basename string `json:"basename"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
)

func NewConfig() *Config {

	value, ok := os.LookupEnv("PG_PORT")
	pgPort, err := strconv.Atoi(value)
	if !ok || err != nil {
		log.Warn("No postgres port passed. Using default 5432 PostgreSQL port")
		// web server local port
		pgPort = 5432
	}
	value, ok = os.LookupEnv("HTTP_PORT")
	cfg := &Config{
		HTTP: HTTP{
			HTTPPort: os.Getenv("HTTP_PORT"),
		},
		PG: PG{
			Address:  os.Getenv("PG_IP"),
			Port:     pgPort,
			Basename: os.Getenv("PG_DATABASE"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
		},
		Log: Log{
			Level: os.Getenv("LOG_LEVEL"),
		},
	}

	return cfg
}

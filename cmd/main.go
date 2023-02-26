package main

import (
	"taskService/config"
	"taskService/internal/app"
)

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}

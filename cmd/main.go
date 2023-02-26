package main

import (
	"taskService/app"

	"taskService/config"
)

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}

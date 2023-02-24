package main

import (
	"fmt"

	"taskService/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Println(cfg)
}

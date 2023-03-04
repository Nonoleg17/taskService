package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"taskService/config"
	"taskService/internal/controller/http"
	"taskService/internal/usecase"
	"taskService/internal/usecase/repo"
	"taskService/pkg/httpserver"
	"taskService/pkg/logger"
	"taskService/pkg/postgres"
	"taskService/pkg/redis"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	pg, err := postgres.New(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	rd, err := redis.New(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - redis.New: %w", err))
	}

	userСase := usecase.NewUserCase(repo.NewUserRepo(pg), repo.NewSessionRepo(rd))

	handler := gin.New()
	http.NewRouter(handler, l, userСase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTPPort))
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

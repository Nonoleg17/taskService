package http

import (
	"taskService/internal/usecase"
	"taskService/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *gin.Engine, l logger.Interface, u usecase.User) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)
	h := handler.Group("/")
	{
		newUserRoutes(h, u, l)
	}

}

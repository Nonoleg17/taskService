package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskService/internal/controller/http/middleware"
	"taskService/internal/entity"
	"taskService/internal/usecase"
	"taskService/pkg/logger"
)

type taskRoutes struct {
	u usecase.Task
	l logger.Interface
}

func newTaskRoutes(handler *gin.RouterGroup, u usecase.User, t usecase.Task, l logger.Interface) {
	r := &taskRoutes{t, l}
	h := handler.Group("/task")
	{
		h.POST("/create", middleware.AuthUser(u), r.Create)
	}
}
func (t *taskRoutes) Create(c *gin.Context) {
	var request *entity.Task
	if err := c.ShouldBindJSON(&request); err != nil {
		t.l.Error(err, "http - Create")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	res, err := t.u.Create(c, request)
	if err != nil {
		t.l.Error(err, "http - Create")
		errorResponse(c, http.StatusInternalServerError, "Create task problem")
		return
	}
	c.JSON(http.StatusOK, res)
}

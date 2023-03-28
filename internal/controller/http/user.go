package http

import (
	"net/http"
	"taskService/internal/entity"
	"taskService/internal/usecase"
	"taskService/pkg/logger"

	"github.com/gin-gonic/gin"
)

type userRoutes struct {
	u usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.User, l logger.Interface) {
	r := &userRoutes{u, l}
	h := handler.Group("/user")
	{
		h.POST("/register", r.Register)
		h.POST("/login", r.Login)
	}
}

func (r *userRoutes) Register(c *gin.Context) {
	var request *entity.CreateUserReq
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - Register")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	res, err := r.u.CreateUser(c.Request.Context(), request)
	if err != nil {
		r.l.Error(err, "http - Register")
		errorResponse(c, http.StatusInternalServerError, "Create user problem")
		return
	}
	c.JSON(http.StatusOK, res)
}
func (r *userRoutes) Login(c *gin.Context) {
	var request *entity.LoginUserReq
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - Login")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	res, session, err := r.u.Login(c.Request.Context(), request)
	if err != nil {
		r.l.Error(err, "http - Login")
		errorResponse(c, http.StatusInternalServerError, "Login problem")
		return
	}
	c.SetCookie("session_token", session.Value, int(session.Expire), "/", "localhost", false, false)
	c.JSON(http.StatusOK, res)

}

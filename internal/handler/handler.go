package handler

import (
	"github.com/gin-gonic/gin"
	"rest_api/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})

	return router
}

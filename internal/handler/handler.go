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

	viruses := router.Group("/viruses")
	{
		viruses.GET("/", h.getAllVirus)
		viruses.POST("/", h.createVirus)
		viruses.DELETE("/:virus_id", h.deleteVirus)
		viruses.GET("/:virus_id", h.getVirusById)
		viruses.PUT("/:virus_id", h.updateVirus)
	}

	sites := router.Group("/sites")
	{
		sites.POST("/", h.createSite)
		sites.GET("/", h.getAllSites)
		sites.GET("/:site_id", h.getSiteById)
		sites.DELETE("/:site_id", h.deleteSite)
		sites.PUT("/:site_id", h.updateSite)
	}

	return router
}

package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "rest_api/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		viruses := api.Group("/viruses", h.userIdentity)
		{
			viruses.GET("/", h.getAllVirus)
			viruses.POST("/", h.createVirus)
			viruses.DELETE("/:virus_id", h.deleteVirus)
			viruses.GET("/:virus_id", h.getVirusById)
			viruses.PUT("/:virus_id", h.updateVirus)
		}

		sites := api.Group("/sites")
		{
			sites.POST("/", h.createSite)
			sites.GET("/", h.getAllSites)
			sites.GET("/:site_id", h.getSiteById)
			sites.DELETE("/:site_id", h.deleteSite)
			sites.PUT("/:site_id", h.updateSite)
		}

		attack := api.Group("/attacks")
		{
			attack.GET("/", h.getAllAttack)
			attack.POST("/", h.createAttack)
			attack.DELETE("/:id", h.deleteAttack)
			attack.GET("/:id", h.getAttackById)
			attack.PUT("/:id", h.updateAttack)
		}
	}

	return router
}

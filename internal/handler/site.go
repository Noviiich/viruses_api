package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	app "rest_api"
	"strconv"
)

func (h *Handler) createSite(c *gin.Context) {
	var input app.Site
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	siteId, err := h.service.Site.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"site_id": siteId,
	})
}

func (h *Handler) getAllSites(c *gin.Context) {
	sites, err := h.service.Site.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sites)
}

func (h *Handler) getSiteById(c *gin.Context) {
	siteId, err := strconv.Atoi(c.Param("site_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	site, err := h.service.Site.GetById(siteId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, site)
}

func (h *Handler) deleteSite(c *gin.Context) {
	siteId, err := strconv.Atoi(c.Param("site_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	err = h.service.Site.Delete(siteId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Удалился сайтик",
	})
}

func (h *Handler) updateSite(c *gin.Context) {
	siteId, err := strconv.Atoi(c.Param("site_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	var input app.SiteUpdate
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Site.Update(siteId, input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

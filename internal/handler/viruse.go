package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	app "rest_api"
	"strconv"
)

func (h *Handler) createVirus(c *gin.Context) {
	var input app.Virus
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	virusID, err := h.service.Virus.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"virus_id": virusID,
	})
}

func (h *Handler) getAllVirus(c *gin.Context) {
	viruses, err := h.service.Virus.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, viruses)
}

func (h *Handler) getVirusById(c *gin.Context) {
	virusId, err := strconv.Atoi(c.Param("virus_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	virus, err := h.service.Virus.GetVirusById(virusId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, virus)
}

func (h *Handler) deleteVirus(c *gin.Context) {
	virusId, err := strconv.Atoi(c.Param("virus_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	err = h.service.Virus.Delete(virusId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Удалился элементик",
	})
}

func (h *Handler) updateVirus(c *gin.Context) {
	virusId, err := strconv.Atoi(c.Param("virus_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	var input app.VirusUpdate
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Virus.Update(virusId, input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

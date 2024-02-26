package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	app "rest_api"
	"strconv"
)

func (h *Handler) createAttack(c *gin.Context) {
	var input app.Attack
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input)

	attackId, err := h.service.Attack.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": attackId,
	})
}

func (h *Handler) getAllAttack(c *gin.Context) {
	attacks, err := h.service.Attack.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attacks)
}

func (h *Handler) getAttackById(c *gin.Context) {
	attackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	attack, err := h.service.Attack.GetById(attackId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attack)
}

func (h *Handler) deleteAttack(c *gin.Context) {
	attackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	err = h.service.Attack.Delete(attackId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Аттака удалилась",
	})
}

func (h *Handler) updateAttack(c *gin.Context) {
	attackId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id params")
		return
	}

	var input app.AttackUpdate
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Attack.Update(attackId, input); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Произошло обновление атаки",
	})
}

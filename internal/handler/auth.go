package handler

import (
	"log"
	"net/http"
	app "rest_api"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Description create account
// @ID create account
// @Tags auth
// @Accept       json
// @Produce      json
// @Param input body app.User true "account info"
// @Success      200  {integer}   integer 1
// @Failure      400  {object}  errorResponse
// @Failure      404  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /auth/sing-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input app.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input app.SignInInput
	log.Print(input)

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

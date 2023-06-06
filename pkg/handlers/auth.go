package handlers

import (
	"awesomeProject5/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input types.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accountId, err := h.service.CreateAccount(types.Account{
		Login:    h.service.GenerateLogin(input.Name),
		Password: h.service.GeneratePassword()},
		input.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	userId, err := h.service.CreateUser(input, accountId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})
}
func (h *Handler) signIn(c *gin.Context) {

}

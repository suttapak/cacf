package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suttapak/cacf/dto"
	"github.com/suttapak/cacf/services"
)

type AuthHandler interface {
	SignIn(c *gin.Context)
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{authService}
}

func (h authHandler) SignIn(c *gin.Context) {
	signInDTO := dto.SignInDTO{}
	if err := c.ShouldBindJSON(&signInDTO); err != nil {
		//TODO : handler error.
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	signInReply, err := h.authService.SignIn(signInDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, signInReply)
}

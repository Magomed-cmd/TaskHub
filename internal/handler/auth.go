package handler

import (
	"TaskHub/internal/service"
	"TaskHub/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	s service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{s: s}
}

func (h *AuthHandler) Login(c *gin.Context) {

	ctx := c.Request.Context()
	var data model.LoginInput
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.s.Login(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"token_type":   "bearer",
		"expires_in":   900,
	})
}

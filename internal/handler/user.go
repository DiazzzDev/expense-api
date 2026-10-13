package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example/expense-api/internal/middleware"
	"example/expense-api/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

type updateMeRequest struct {
	DisplayName string `json:"display_name" binding:"required"`
}

func (h *UserHandler) Me(c *gin.Context) {
	u, err := h.svc.Get(c.Request.Context(), c.GetString(middleware.UserIDKey))
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h *UserHandler) UpdateMe(c *gin.Context) {
	var req updateMeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.svc.UpdateDisplayName(c.Request.Context(), c.GetString(middleware.UserIDKey), req.DisplayName)
	if err != nil {
		respondError(c, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

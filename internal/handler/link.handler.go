package handler

import (
	"context"
	"net/http"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/dto"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/gin-gonic/gin"
)

type ILinkService interface {
	GetMyLinks(ctx context.Context, userID int) ([]model.ResLink, error)
}

type LinkHandler struct {
	Service ILinkService
}

func NewLinkHandler(s ILinkService) *LinkHandler {
	return &LinkHandler{Service: s}
}

func (h *LinkHandler) GetMyLinks(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "Unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "Invalid user identity",
		})
		return
	}

	links, err := h.Service.GetMyLinks(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "Failed to fetch links",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "Links retrieved successfully",
		Data:    links,
	})
}

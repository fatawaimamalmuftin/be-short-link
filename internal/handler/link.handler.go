package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/dto"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type ILinkService interface {
	CreateShortLink(ctx context.Context, userID int, req *model.ReqCreateLink) (*model.ResLink, error)
	GetMyLinks(ctx context.Context, userID int) ([]model.ResLink, error)
	DeleteShortLink(ctx context.Context, linkID int, userID int) error
}

type LinkHandler struct {
	Service ILinkService
}

func NewLinkHandler(s ILinkService) *LinkHandler {
	return &LinkHandler{Service: s}
}

func (h *LinkHandler) CreateLink(c *gin.Context) {
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

	var req model.ReqCreateLink
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: err.Error(),
		})
		log.Println("Validation Error CreateLink:", err)
		return
	}

	link, err := h.Service.CreateShortLink(c.Request.Context(), userID, &req)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, dto.Res{
				Status:  false,
				Message: "Slug already exists, please choose another slug",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "Failed to create short link",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Res{
		Status:  true,
		Message: "Short link created successfully",
		Data:    link,
	})
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

func (h *LinkHandler) DeleteLink(c *gin.Context) {
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

	idParam := c.Param("id")
	linkID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: "Invalid link ID parameter",
		})
		return
	}

	if err := h.Service.DeleteShortLink(c.Request.Context(), linkID, userID); err != nil {
		c.JSON(http.StatusNotFound, dto.Res{
			Status:  false,
			Message: "Link not found or not owned by user",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "Short link deleted successfully",
	})
}

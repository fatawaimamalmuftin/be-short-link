package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	customerror "github.com/fatawaimamalmuftin/be-shortlink/internal/customError"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/dto"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type IAuthService interface {
	RegisterUser(ctx context.Context, newUser *model.ReqRegis) error
	LoginUser(ctx context.Context, req *model.ReqLogin) (*model.ResLogin, error)
	LogoutUser(ctx context.Context, token string) error
}

type AuthHandler struct {
	Service IAuthService
}

func NewAuthHandler(s IAuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req model.ReqRegis

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: err.Error(),
		})
		log.Println("Validation Error:", err)
		return
	}

	if err := h.Service.RegisterUser(c.Request.Context(), &req); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, dto.Res{
				Status:  false,
				Message: "Email is already registered",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "Failed to process registration",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Res{
		Status:  true,
		Message: "User registered successfully",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.ReqLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: err.Error(),
		})
		log.Println("Validation Error:", err)
		return
	}

	res, err := h.Service.LoginUser(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, customerror.PasMissMach) {
			c.JSON(http.StatusUnauthorized, dto.Res{
				Status:  false,
				Message: "Invalid email or password",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "Failed to process login",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "Login successful",
		Data:    res,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: "Authorization header is required",
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, dto.Res{
			Status:  false,
			Message: "Invalid token format",
		})
		return
	}

	token := parts[1]

	if err := h.Service.LogoutUser(c.Request.Context(), token); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "Failed to process logout",
		})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.Res{
		Status:  true,
		Message: "Logout successful",
	})
}

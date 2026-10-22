package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/dto"
	jwtpkg "github.com/fatawaimamalmuftin/be-shortlink/pkg"
	"github.com/fatawaimamalmuftin/be-shortlink/pkg/blacklist"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func CheckJWTtoken(c *gin.Context, rdb *redis.Client) {
	autorizationhHeader := c.GetHeader("Authorization")

	if autorizationhHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "please login first",
		})
		return
	}

	bearer := strings.Split(autorizationhHeader, " ")

	if len(bearer) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "invalid bearer token",
		})
		return
	}

	if bearer[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "invalid bearer token",
		})
		return
	}

	var tokenClem jwtpkg.JWTclem

	err := tokenClem.DecodeToken(bearer[1])

	if err != nil {
		if errors.Is(err, jwt.ErrTokenInvalidIssuer) || errors.Is(err, jwt.ErrTokenExpired) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Res{
				Status:  false,
				Message: "invalid token",
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "internal server error",
		})
		log.Println(err)
		return
	}

	isBlacklisted, err := blacklist.IsTokenBlackList(bearer[1], rdb, c.Request.Context())

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Res{
			Status:  false,
			Message: "internal server error",
		})
		log.Println(err)
		return
	}

	if isBlacklisted {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Res{
			Status:  false,
			Message: "invalid token",
		})
		return
	}

	c.Set("tokenCleims", tokenClem)
	c.Set("user_id", tokenClem.Id)
	c.Next()
}

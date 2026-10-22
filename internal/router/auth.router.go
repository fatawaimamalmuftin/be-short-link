package router

import (
	"context"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/handler"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/repo"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func authRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client, c *context.Context) {
	auth := r.Group("/api/auth")

	rr := repo.NewAuthRepo(db)
	rs := service.NewAuthService(rr)
	rh := handler.NewAuthHandler(rs)
	auth.POST("registration", rh.Register)
}

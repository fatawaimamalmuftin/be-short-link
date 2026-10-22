package router

import (
	"context"

	middleware "github.com/fatawaimamalmuftin/be-shortlink/internal/Middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func MainRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client, c *context.Context) {
	r.Use(middleware.Cors)
	authRouter(r, db, rdb, c)
}

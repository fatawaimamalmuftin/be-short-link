package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func authRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client, c *context.Context) {

}

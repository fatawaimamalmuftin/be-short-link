package router

import (
	Middleware "github.com/fatawaimamalmuftin/be-shortlink/internal/Middleware"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/handler"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/repo"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func LinkRouter(r *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	linkRepo := repo.NewLinkRepo(db)
	linkService := service.NewLinkService(linkRepo)
	linkHandler := handler.NewLinkHandler(linkService)

	linkGroup := r.Group("/api/links")
	linkGroup.Use(func(c *gin.Context) {
		Middleware.CheckJWTtoken(c, rdb)
	})
	{
		linkGroup.GET("/my", linkHandler.GetMyLinks)
	}
}

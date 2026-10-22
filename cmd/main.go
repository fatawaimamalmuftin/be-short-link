package main

import (
	"context"
	"log"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/config/database"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/config/redis"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load .env")
	}

	r := gin.Default()
	r.Static("/uploads", "./public/uploads")

	db := database.InitDB()

	rdb := redis.InitRedis()

	ctx := context.Background()
	router.MainRouter(r, db, rdb, &ctx)

	r.Run(":3000")
}

package main

import (
	"context"
	"log"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/config/database"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/config/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load .env")
	}

	r := gin.Default()
	r.Static("/uploads", "./public/uploads")

	if err := database.InitDB(); err != nil {
		log.Println(err)
	}

	if err := redis.InitRedis(); err != nil {
		log.Println(err)
	}

	ctx := context.Background()

	r.Run(":3000")
}

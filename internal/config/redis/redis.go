package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rdbHost := os.Getenv("RDB_HOST")
	rdbPort := os.Getenv("RDB_PORT_H")

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rdbHost, rdbPort),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Println("Unable to connect to Redis:", err)
	}

	log.Println("Connected to Redis successfully")
	return rdb
}

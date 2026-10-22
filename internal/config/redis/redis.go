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

	if rdbHost == "" {
		rdbHost = "localhost"
	}
	if rdbPort == "" {
		rdbPort = "8002"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rdbHost, rdbPort),
		Username: "shortlink",
		Password: "123",
		DB:       0,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Unable to connect to Redis: %v\n", err)
	}

	log.Println("Connected to Redis successfully")
	return rdb
}

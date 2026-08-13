package db

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitializeRedisDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 2,
	})
	ctx := context.Background()

	err := rdb.Set(ctx, "salam", "balapan", 10*time.Second).Err()
	if err != nil {
		log.Fatal(err)
	}
}

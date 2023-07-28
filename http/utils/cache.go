package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-micro.dev/v4/util/log"
)

func NewRedisClient(ctx context.Context) (*redis.Client, func()) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// 连接 Redis
	if err := client.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	log.Info("connect redis success")
	return client, func() {
		client.Close()
	}
}

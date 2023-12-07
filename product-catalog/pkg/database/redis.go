package database

import (
	"context"
	"product-catalog/config"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis(cfg config.Redis) (client *redis.Client, err error) {
	ctx := context.Background()
	// address := cfg.Host + ":" + cfg.Port
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
	})

	err = client.Ping(ctx).Err()
	if err != nil {
		return
	}
	return
}

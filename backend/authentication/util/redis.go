package util

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

func NewRedis(db int) *redis.Client {
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       db,
	})

	return redisClient
}

func GetRedis(rdb *redis.Client, key string) (interface{}, error) {
	ctx := context.Background()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func SetRedis(rdb *redis.Client, key string, value interface{}, expTime time.Duration) error {
	ctx := context.Background()
	err := rdb.Set(ctx, key, value, expTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func DeleteRedis(rdb *redis.Client, key string) error {
	ctx := context.Background()
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

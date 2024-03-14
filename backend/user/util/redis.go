package util

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "103.184.113.181:6379",
	Password: "nqm05112000",
	DB:       1,
})

var ctx = context.Background()

//set to redis
func SetRedis(key string, value interface{}, expTime time.Duration) error {
	err := rdb.Set(ctx, key, value, expTime).Err()
	if err != nil {
		return err
	}

	return nil
}

//get redis
func GetRedis(key string) (interface{}, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func DeleteRedis(key string) error {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

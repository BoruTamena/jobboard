package rediscache

import (
	"context"
	"time"

	"github.com/BoruTamena/jobboard/internal/constant/models/dto"
	"github.com/BoruTamena/jobboard/platform"
	"github.com/go-redis/redis/v8"
)

type redisGo struct {
	client *redis.Client
}

func InitRedisCache(cfg dto.Config) platform.RedisCache {

	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host + ":" + cfg.Redis.Port,
	})

	return &redisGo{
		client: client,
	}

}

func (cache *redisGo) AddValue(ctx context.Context, key string, value interface{}, exp int64) error {

	duration := time.Until(time.Unix(exp, 0))

	err := cache.client.Set(ctx, key, value, duration).Err()

	return err

}

func (cache *redisGo) GetValue(ctx context.Context, key string) (interface{}, error) {

	var val interface{}

	err := cache.client.Get(ctx, key).Scan(val)

	return val, err

}

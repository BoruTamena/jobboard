package platform

import "context"

type RedisCache interface {
	AddValue(ctx context.Context, key string, value interface{}, exp int64) error
	GetValue(ctx context.Context, key string) (interface{}, error)
}

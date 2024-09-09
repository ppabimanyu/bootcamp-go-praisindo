package redisser

import (
	"context"
	"time"
)

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	HGet(ctx context.Context, key, field string) (string, error)
	HSet(ctx context.Context, key, field, value string) error
	HDel(ctx context.Context, key, field string) error
	SetHashesExpire(ctx context.Context, key string, timeopen, timeclose int64) error
	SetWithExpire(ctx context.Context, key string, value interface{}, second time.Duration) (string, error)
	Set(ctx context.Context, key string, value interface{}) (string, error)
	Del(ctx context.Context, key string) (int64, error)
	SetBit(ctx context.Context, key string, offset int64, value int) (int64, error)
	GetAllBits(ctx context.Context, key string) ([]bool, error)
	GetTTL(ctx context.Context, key string) int
}

package database

import (
	"ShawnOJ/utils/setting"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedis(cfg *setting.RedisConfig) (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = Rdb.Ping(context.Background()).Result()
	return
}

func CloseRedis() {
	_ = Rdb.Close()
}

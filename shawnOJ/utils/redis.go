package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

// 创建 Redis 客户端
var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis 地址
		Password: "",               // 无密码可留空
		DB:       0,                // 默认数据库
	})

	// 测试连接
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}
	fmt.Println("Successfully connected to Redis")
}

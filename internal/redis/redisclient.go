package redis

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var Client *redis.Client

func InitRedis() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	Client = redis.NewClient(&redis.Options{
		Addr : addr,})
	}

func GetRedisClient() *redis.Client {
	return Client
}

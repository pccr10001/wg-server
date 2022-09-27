package device

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

var RedisClient *redis.Client

func InitRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Network: "unix",
		Addr:    "/var/run/redis/redis-server.sock",
	})
	if RedisClient.Ping(context.Background()).String() == "ping: PONG" {
		fmt.Fprintf(os.Stderr, "Redis connected.\n")
	}
}

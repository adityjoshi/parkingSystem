package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

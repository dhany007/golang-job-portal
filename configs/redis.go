package configs

import (
	"github.com/dhany007/golang-job-portal/services/utils"
	redis "github.com/redis/go-redis/v9"
)

func (c *Config) InitRedis() error {
	c.RedisClient = NewRedis()
	return nil
}

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Network: utils.GetEnv("REDIS_CACHE_URL", REDIS_CACHE_URL),
		Addr:    utils.GetEnv("REDIS_CACHE_URL", ""),
	})
}

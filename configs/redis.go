package configs

import (
	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services/utils"
	redis "github.com/redis/go-redis/v9"
)

func (c *Config) InitRedis() error {
	c.RedisClient = NewRedis()
	return nil
}

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     utils.GetEnv("REDIS_CACHE_URL", models.REDIS_CACHE_URL),
		Password: utils.GetEnv("REDIS_CACHE_PWD", ""),
		DB:       0,
	})
}

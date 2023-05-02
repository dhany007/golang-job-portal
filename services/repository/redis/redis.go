package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/utils"
	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	redisClient *redis.Client
}

func NewCacheRepository(redisClient *redis.Client) services.CacheRepository {
	return cacheRepository{redisClient}
}

// GenerateCacheKey implements services.CacheRepository
func (c cacheRepository) GenerateCacheKey(ctx context.Context, prefix models.CachePrefix, name string, parameters ...interface{}) (res string) {
	var parameterStr string
	for i, v := range parameters {
		if i > 0 {
			parameterStr += ", "
		}
		parameterStr += utils.InterfaceToString(v)
	}

	return fmt.Sprintf("%s %s(%s)", prefix, name, utils.SHA1(parameterStr))
}

// Get implements services.CacheRepository
func (c cacheRepository) Get(ctx context.Context, key string) (data string) {
	if c.redisClient == nil {
		log.Println("redis client not ready")
		return
	}

	data, err := c.redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Printf("error while redis get operation key: %s, err:%s\n", key, err.Error())
		return
	}

	return
}

// Set implements services.CacheRepository
func (c cacheRepository) Set(ctx context.Context, key string, value interface{}, duration time.Duration) (err error) {
	if c.redisClient == nil {
		return errors.New("redis client not ready")
	}

	return c.redisClient.Set(ctx, key, value, duration).Err()
}

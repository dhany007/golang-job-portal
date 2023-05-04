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

func (c cacheRepository) DeleteByPrefix(ctx context.Context, prefix models.CachePrefix) (err error) {
	var (
		allKeys []string
		cursor  uint64
	)
	for {
		var (
			keys []string
		)

		keys, cursor, err = c.redisClient.Scan(ctx, cursor, fmt.Sprintf("%s *", prefix), 0).Result()

		if err != nil {
			err = fmt.Errorf("failed to get the redis keys (%s)", prefix)
			return
		}

		allKeys = append(allKeys, keys...)

		if cursor == 0 {
			break
		}
	}

	err = c.redisClient.Del(ctx, allKeys...).Err()
	if err != nil {
		err = fmt.Errorf("failed to delete keys redis (prefix: %s)", prefix)
		return
	}

	return
}

func (c cacheRepository) DeleteByPrefixAsync(ctx context.Context, prefix models.CachePrefix) {
	go func() {
		err := c.DeleteByPrefix(ctx, prefix)
		if err != nil {
			log.Printf("failed to delete keys redis (prefix: %s)\n", prefix)
		}
	}()
}

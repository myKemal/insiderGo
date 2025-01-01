package services

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

var redisCtx = context.Background()

func NewRedisClient(addr string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisClient{Client: client}
}

func (r *RedisClient) Set(key string, value string, expiration time.Duration) error {
	return r.Client.Set(redisCtx, key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.Client.Get(redisCtx, key).Result()
}

func (r *RedisClient) List(skip, limit int) ([]map[string]interface{}, error) {
	keys, err := r.Client.Keys(redisCtx, "*").Result()
	if err != nil {
		return nil, err
	}

	if skip > len(keys) {
		return []map[string]interface{}{}, nil
	}
	if skip+limit > len(keys) {
		limit = len(keys) - skip
	}
	keys = keys[skip : skip+limit]

	var result []map[string]interface{}
	for _, key := range keys {
		value, err := r.Client.Get(redisCtx, key).Result()
		if err != nil {
			continue
		}

		result = append(result, map[string]interface{}{
			"key":   key,
			"value": value,
		})
	}

	return result, nil
}

func (r *RedisClient) AllList() ([]map[string]interface{}, error) {
	keys, err := r.Client.Keys(redisCtx, "*").Result()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, key := range keys {
		value, err := r.Client.Get(redisCtx, key).Result()
		if err != nil {
			continue
		}

		result = append(result, map[string]interface{}{
			"key":   key,
			"value": value,
		})
	}

	return result, nil
}

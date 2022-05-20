package redis

import (
	"context"
	"fmt"
	"myapp/core/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	keyPrefix string
	driver    *redis.Client
}

func NewRedisClient() *RedisClient {
	host := config.GetString("redis.host")
	port := config.GetInt("redis.port")

	driver := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		DB:   config.GetInt("redis.db"),
	})

	return &RedisClient{
		keyPrefix: config.GetString("redis.key_prefix"),
		driver:    driver,
	}
}

func (c *RedisClient) castkey(key string) string {
	return fmt.Sprintf("%s%s", c.keyPrefix, key)
}

func (c *RedisClient) get(key string) *redis.StringCmd {
	return c.driver.Get(context.Background(), c.castkey(key))
}

func (c *RedisClient) GetString(key string) string {
	return c.get(key).String()
}

func (c *RedisClient) GetInt(key string) int {
	// 忽略err处理
	v, _ := c.get(key).Int()
	return v
}

func (c *RedisClient) Set(key string, value interface{}) error {
	return c.SetEx(key, value, 0)
}

func (c *RedisClient) SetEx(key string, value interface{}, ttl time.Duration) error {
	return c.driver.Set(context.Background(), c.castkey(key), value, ttl).Err()
}

func (c *RedisClient) SetNx(key string, value interface{}) (bool, error) {
	return c.SetNxEx(key, value, 0)
}

func (c *RedisClient) SetNxEx(key string, value interface{}, ttl time.Duration) (bool, error) {
	cmd := c.driver.SetNX(context.Background(), key, value, ttl)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}
	return cmd.Result()
}

func (c *RedisClient) Del(keys ...string) error {
	if len(keys) == 0 {
		return fmt.Errorf("keys length must be greater than 0 ")
	}
	return c.driver.Del(context.Background(), keys...).Err()
}

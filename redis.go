package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() (*RedisClient, error) {
	slog.Info("Redis connection:", "string", os.Getenv("REDIS_URL"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"), // Redis server address
		Password: "",                     // No password set
		DB:       0,                      // Use default DB
	})

	// Ping the Redis server to check if it's up and running
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		slog.Error("Error connecting to Redis", "error", err)
		return nil, err
	}
	slog.Info("Connected to Redis")
	return &RedisClient{client: rdb}, nil
}

func (rdb *RedisClient) WriteValueWithTTL(key string, value any, ttl time.Duration) error {
	ctx := context.Background()
	err := rdb.client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("error setting value: %w", err)
	}
	return nil
}

func (rdb *RedisClient) ReadValue(key string) (string, error) {
	ctx := context.Background()
	val, err := rdb.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("error getting value: %w", err)
	}
	return val, nil
}

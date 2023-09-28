package store

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shoet/go-cli-with-kvs/config"
)

type KVSRedis struct {
	redisClient *redis.Client
	experiation time.Duration
}

func NewRedisKVS(cfg *config.Config) (*KVSRedis, error) {
	cli := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%d", cfg.KVSHost, cfg.KVSPort),
		Username: cfg.KVSUser,
		Password: cfg.KVSPassword,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	})
	status := cli.Ping(context.Background())
	if status.Err() != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", status.Err())
	}
	kvs := &KVSRedis{
		redisClient: cli,
		experiation: time.Duration(cfg.KVSExperiationSec) * time.Second,
	}

	return kvs, nil
}

func (kvs *KVSRedis) Get(ctx context.Context, key string) (string, error) {
	val := kvs.redisClient.Get(ctx, key)
	return val.Val(), val.Err()
}

func (kvs *KVSRedis) Set(ctx context.Context, key string, value interface{}) error {
	return kvs.redisClient.Set(ctx, key, value, kvs.experiation).Err()
}

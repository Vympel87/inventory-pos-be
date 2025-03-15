package config

import (
	"context"
	"fmt"

	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	url := fmt.Sprintf("redis://default:%s@%s:%s", Config.RedisPass, Config.RedisHost, Config.RedisPort)

	opt, err := redis.ParseURL(url)
	if err != nil {
		Log.WithField(enum.LevelError, err).Fatal("failed to parse redis URL")
	}

	RedisClient = redis.NewClient(opt)

	_, err = RedisClient.Ping(context.Background()).Result()
	if err != nil {
		Log.WithField("error", err).Fatal("failed to connect to redis")
	}

	Log.WithFields(map[string]interface{}{
		enum.LevelInfo: "Redis connection established",
		"url":          url,
	}).Info("Connected to Redis")
}

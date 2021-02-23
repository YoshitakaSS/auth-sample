package infra

import (
	"fmt"

	"github.com/YoshitakaSS/go_auth/config"
	"github.com/gomodule/redigo/redis"
)

func initRedis() (redis.Conn, error) {
	redis, err := redis.Dial("tcp", config.RedisDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect redis %w", err)
	}

	return redis, nil
}

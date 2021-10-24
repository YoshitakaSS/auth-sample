package infra

import (
	"fmt"

	"github.com/YoshitakaSS/go_auth/config"
	"github.com/gomodule/redigo/redis"
)

// InitRedis Connect Redis and Return Redis Connect Info
func InitRedis() (redis.Conn, error) {
	redis, err := redis.Dial("tcp", config.RedisDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect redis %w", err)
	}

	return redis, nil
}

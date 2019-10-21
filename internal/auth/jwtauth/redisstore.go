package jwtauth

import (
	"time"

	"github.com/go-redis/redis"
)

//RedisConfig redis配置参数
type RedisConfig struct {
	addr      string
	password  string
	keyprefix string
}

type RedisStore struct {
	cli       *redis.Client
	keyPrefix string
}

func NewStore(cfg *RedisConfig) *RedisStore {
	rediscli := redis.NewClient(&redis.Options{
		Addr:     cfg.addr,
		Password: cfg.password,
	})

	return &RedisStore{
		cli:       rediscli,
		keyPrefix: cfg.keyprefix,
	}
}

func (s *RedisStore) Set(token string, expire time.Duration) error {
	cmd := s.cli.Set(token, "jwttoken", expire)
	return cmd.Err()
}

func (s *RedisStore) Verify(token string) (bool, error) {
	cmd := s.cli.Exists(token)
	if cmd.Err() != nil {
		return false, cmd.Err()
	}
	return cmd.Val() > 0, nil
}

func (s *RedisStore) Close() {
	s.cli.Close()
}

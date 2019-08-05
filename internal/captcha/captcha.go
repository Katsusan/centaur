package captcha

import (
	"encoding/hex"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type redisStore struct {
	cli        *redis.Client
	prefix     string
	expiration time.Duration
}

func (s *redisStore) keyWithPrefix(key string) string {
	return s.prefix + key
}

func (s *redisStore) Set(id string, digits []byte) {
	if s.cli == nil {
		log.Println("redis client not established")
		return
	}

	exec := s.cli.Set(s.keyWithPrefix(id), hex.EncodeToString(digits), s.expiration)
	if err := exec.Err(); err != nil {
		log.Printf("failed to exec redis set command [%v]", err)
	}
	return
}

func (s *redisStore) Get(id string, clear bool) (digits []byte) {
	exec := s.cli.Get(s.keyWithPrefix(id))
	if err := exec.Err(); err != nil {
		if err == redis.Nil {
			return nil
		}
		log.Printf("failed to exec redis get command [%v]", err)
		return nil
	}

	res, err := hex.DecodeString(exec.Val())
	if err != nil {
		log.Printf("failed to decode result from redis get command [%v]", err)
	}

	//determine if delete this key
	if clear {
		exec := s.cli.Del(s.keyWithPrefix(id))
		if err := exec.Err(); err != nil {
			log.Printf("failed to exec redis del command [%v]", err)
			return nil
		}

	}

	return res
}

func NewRedisStore(opts *redis.Options, prefix string, expiration time.Duration) *redisStore {
	if opts == nil {
		log.Fatalln("cannot initialize RedisStore with null options")
	}

	return &redisStore{
		cli:        redis.NewClient(opts),
		prefix:     prefix,
		expiration: expiration,
	}
}

func InitCaptcha() {

}

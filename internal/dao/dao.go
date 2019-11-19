package dao

import (
	"github.com/jinzhu/gorm"
	"log"

	"github.com/go-redis/redis"
)

//Dao struct info of database access object
type Dao struct {
	db    *gorm.DB
	redis *redis.Client
}

func New() *Dao {
	dbconn, err := gorm.Open("")
	if err != nil {
		log.Panicln("failed to connect to DB", err)
	}

	return &Dao{
		db:    dbconn,
		redis: redis.NewClient(&redis.Options{}),
	}
}

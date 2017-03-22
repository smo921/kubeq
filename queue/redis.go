package queue

import (
	"time"

	"github.com/go-redis/redis"
)

func DoRedisStuff() (result string, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          0,  // use default DB
		DialTimeout: time.Second * 2,
	})

	result, err = client.LPop("foobar").Result()
	client.Close()
	return
}

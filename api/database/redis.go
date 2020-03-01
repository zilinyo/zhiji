package database

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//defer Client.Close()

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
		panic("redis not connect")
	}
	fmt.Printf("ping result: %s\n", pong)

}

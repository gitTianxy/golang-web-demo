package base

import (
	"github.com/go-redis/redis"
	"fmt"
)

func GetRedisClient() *redis.Client {
	conf := GetRedisConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := client.Ping().Result()
	CheckErr(err)
	return client
}

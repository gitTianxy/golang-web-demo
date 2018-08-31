package rao

import (
	"github.com/go-redis/redis"
	"time"
	"golang-web-demo/base"
)

const EXPIRE_SECS = 60

type ItemRao struct {
	client *redis.Client
}

func (rao *ItemRao) SetClient(client *redis.Client)  {
	rao.client = client
}

func (rao ItemRao) Set(key string, val string) {
	_, err := rao.client.Set(key, val, EXPIRE_SECS * time.Second).Result()
	base.CheckErr(err)
}

func (rao ItemRao) Get(key string) (val string, ok bool) {
	ok = true
	val, err := rao.client.Get(key).Result()
	if err != nil {
		ok = false
	}
	return
}


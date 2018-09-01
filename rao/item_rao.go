package rao

import (
	"github.com/go-redis/redis"
	"time"
	"golang-web-demo/model"
	"encoding/json"
	"strings"
	"golang-web-demo/util"
)

const EXPIRE_SECS = 300
const KEY_TMPL = "item_{id}"

type ItemRao struct {
	client *redis.Client
}

func (rao *ItemRao) SetClient(client *redis.Client)  {
	rao.client = client
}

func (rao ItemRao) getKey(id int64) string  {
	return strings.Replace(KEY_TMPL, "{id}", util.Int642String(id), 1)
}

func (rao ItemRao) getExpiration() time.Duration {
	return EXPIRE_SECS * time.Second
}

func (rao ItemRao) Set(item model.Item) error {
	val, err := json.Marshal(item)
	if err != nil {
		return err
	}
	key := rao.getKey(item.ID)
	expire := rao.getExpiration()
	return rao.client.Set(key, val, expire).Err()
}

func (rao ItemRao) Get(id int64) (item model.Item, err error) {
	key := rao.getKey(id)
	val, err := rao.client.Get(key).Result()
	// ignore redis.Nil
	if err == redis.Nil {
		err = nil
	}
	if len(val) > 0 {
		json.Unmarshal([]byte(val), &item)
	}
	return
}

func (rao ItemRao) Del(id int64) error {
	key := rao.getKey(id)
	return rao.client.Del(key).Err()
}


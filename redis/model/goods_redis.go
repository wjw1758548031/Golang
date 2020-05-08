package model

import (
	"fmt"
)

type GoodsRedis struct {
	RedisModel
}

const (
	ID = "goods_"
)

func (this *GoodsRedis) Getstring(id string) string {
	key := fmt.Sprintf(ID + id)
	return Pool.Getstrings(key)
}

func (this *GoodsRedis) Setstring(id string, value interface{}) {
	key := fmt.Sprintf(ID + id)
	Pool.Setstrings(key, value)
}

package controllers

import (
	"log"
	"redis/model"
)

type Redis struct {
	ApiController
}

// @Description wenjianliu 互斥所，锁住所有携程*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /getRedis [post]
func (this *Redis) GetRedis() {
	log.Print("进入GetRedis")
	goodsRedis := model.GoodsRedis{}
	item := goodsRedis.Getstring("1")
	log.Print("GetRedis:", item)
}

// @Description wenjianliu 互斥所，锁住所有携程*******************************
// @Param body body controllers.Zhi true "请求参数说明"
// @Success 200 {object} controllers.Zhi
// @Failure 201 {"err_code":"错误编码值","err_msg":"错误描述值"}
// @router /setRedis [post]
func (this *Redis) SetRedis() {
	log.Print("进入SetRedis")
	goodsRedis := model.GoodsRedis{}
	goodsRedis.Setstring("1", "王建文")
}

package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/lunny/log"
	"time"
)

type RedisModel struct {
	redisPool *redis.Pool
}

var Pool RedisModel

func (this *RedisModel) Init() {
	log.Println("进入redis连接")
	//"tcp", "127.0.0.1:6379", "", 0
	network := "tcp"
	address := "127.0.0.1:6379" //开启就能看到端口
	password := "123456"        //redis.conf 里去查看密码和修改密码 搜索 requirepass
	this.redisPool = &redis.Pool{
		MaxIdle:     100,               //最大空闲连接数
		IdleTimeout: 300 * time.Second, //	    // 空闲连接超时时间，超过超时时间的空闲连接会被关闭。 如果设置成0，空闲连接将不会被关闭 应该设置一个比redis服务端超时时间更短的时间
		Dial: func() (redis.Conn, error) { //Dial()方法返回一个连接，从在需要创建连接到的时候调用
			opts := []redis.DialOption{redis.DialDatabase(0)}
			if len(password) > 0 {
				opts = append(opts, redis.DialPassword(password))
			}
			c, err := redis.Dial(network, address, opts...)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				log.Println("redis连接失败:", err)
			}
			return err
		},
	}

	//pool = this.redisPool
	/*	c := this.redisPool.Get()
		defer c.Close()
		_,err := c.Do("PING");if err!=nil{
			log.Println("redis连接失败:",err)
		}*/
	log.Println("reds连接成功")
	return
}

func (this *RedisModel) Getstrings(id string) string {
	c := this.redisPool.Get()
	defer c.Close()
	reply, err := c.Do("GET", id)
	if err != nil {
		fmt.Println("redis-get()错误:", err)
	}
	item, err := redis.String(reply, err)
	if err != nil {
		fmt.Println("redis-get()错误:", err)
	}
	//redis.Bytes()
	return item
}

func (this *RedisModel) Setstrings(id string, value interface{}) {
	c := this.redisPool.Get()
	defer c.Close()
	_, err := c.Do("SET", id, value)
	if err != nil {
		fmt.Println("redis-SET()错误:", err)
	}
}

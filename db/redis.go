package db

import (
	"github.com/astaxie/beego"
	redis "github.com/yangfei4913438/redis-full"
	"time"
)

/*
# redis配置
redis.host=10.0.0.253:6379
redis.password=
redis.db=0
redis.maxidle=100
redis.maxactive=1000
redis.idletimeout=600
*/

//redis对外接口
var RedisDB redis.RedisCache

func init() {
	initRedis()
}

func initRedis() {
	hosts := beego.AppConfig.String("redis.host")
	password := beego.AppConfig.DefaultString("redis.password", "")
	database := beego.AppConfig.DefaultInt("redis.db", 0)
	MaxIdle := beego.AppConfig.DefaultInt("redis.maxidle", 100)
	MaxActive := beego.AppConfig.DefaultInt("redis.maxactive", 1000)
	IdleTimeout := time.Duration(beego.AppConfig.DefaultInt("redis.idletimeout", 600)) * time.Second

	//通过赋值对外接口来使用
	RedisDB = redis.NewRedisCache(hosts, password, database, MaxIdle, MaxActive, IdleTimeout, 24*time.Hour)

	if err := RedisDB.CheckRedis(); err != nil {
		panic("Redis Server:" + hosts + " Connect failed: " + err.Error() + "!")
	} else {
		beego.Info("Redis Server:" + hosts + " Connected!")
	}
}

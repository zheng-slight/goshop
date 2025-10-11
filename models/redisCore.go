package models

//redis官网: github.com/go-redis
//下载go-redis: go get github.com/redis/go-redis/v9
//连接redis数据库核心代码

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
	"os"
)

//全局使用,就需要把定义成公有的
var ctxRedis = context.Background()

var (
	RedisDb *redis.Client
)

//是否开启redis
var redisEnable bool

//自动初始化数据库
func init() {
	//加载配置文件
	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	//获取redis配置
	ip := config.Section("redis").Key("ip").String()
	port := config.Section("redis").Key("port").String()
	redisEnable, _ = config.Section("redis").Key("redisEnable").Bool()

	//判断是否开启redis
	if redisEnable {
		RedisDb = redis.NewClient(&redis.Options{
			Addr:     ip + ":" + port,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		//连接redis
		_, err := RedisDb.Ping(ctxRedis).Result()
		//判断连接是否成功
		if err != nil {
			println(err)
		}
	}
}
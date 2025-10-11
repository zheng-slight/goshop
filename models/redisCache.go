package models

import (
	"encoding/json"
	"time"
)

type RedisCache struct {
}

//设置
func (r RedisCache) Set(key string, value interface{}, expiration int) {
	if redisEnable {  //判断是否开启redis
		v, err := json.Marshal(value)  //value是一个空接口类型,里面可以是字符串,切片,结构体,所以转成json保存
		if err == nil {
			RedisDb.Set(ctxRedis, key, string(v), time.Second*time.Duration(expiration))
		}
	}
}

//获取
func (r RedisCache) Get(key string, obj interface{}) bool {
	if redisEnable { //判断是否开启redis
		valueStr, err1 := RedisDb.Get(ctxRedis, key).Result()
		if err1 == nil && valueStr != "" {
			err2 := json.Unmarshal([]byte(valueStr), obj)
			return err2 == nil
		}
	}
	return false
}

//清除缓存
func (r RedisCache) FlushAll() {
	if redisEnable {
		RedisDb.FlushAll(ctxRedis)
	}
}

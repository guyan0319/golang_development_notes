package main

import (
	"example/example/public/redispool"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var RedisExpire = 3600 //缓存有效期
func main() {

	// 从池里获取连接
	rc := redispool.RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()
	key := "redis.cache"
	_, err := rc.Do("Set", key, "1", "EX", RedisExpire)
	if err != nil {
		fmt.Println(err)
		return
	}
	val, err := redis.String(rc.Do("Get", key))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
	//删除
	rc.Do("Del", key)

}

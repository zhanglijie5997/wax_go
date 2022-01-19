package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	os2 "go_study/utils/os"
	"time"
)
var ctx = context.Background()
func ClientRedis()  {
	_yaml := os2.YamlResult.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:		_yaml.Address,
		Password:   _yaml.Password,
		DB: 		0,
	})
	err := rdb.Set(ctx, "key", "123", 0).Err()
	_err := rdb.SetNX(ctx, "key2", "12344", 10 * time.Second).Err()
	if err != nil && _err != nil {
		panic(err)
	}else  {
		fmt.Println("redis connect success!!")
	}
	val, err := rdb.Get(ctx, "key").Result()
	key2, _key2Err := rdb.Get(ctx, "key2").Result()
	if err == nil  && _key2Err == nil{
		fmt.Println(val, key2,"获取redis 数据 -->>> ")
	}
}
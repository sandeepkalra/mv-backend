package utils

import (
	"fmt"
	"time"

	redis "gopkg.in/redis.v3"
)

type RedisDb struct {
	R *redis.Client
}

func FastMemInit(redis_ip string) (bool, *RedisDb) {
	dial := redis_ip + ":6379"
	client := redis.NewClient(&redis.Options{
		Addr:     dial,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis err", pong, err)
		return false, nil
	}
	redis := RedisDb{}
	redis.R = client
	return true, &redis
}

func (obj *RedisDb) FastMemDestroy() {
	obj.R.Close()
}

func (myredis_client *RedisDb) TimedAdd(stream string, key string, val string) bool {
	new_key := key + stream
	myredis_client.R.Set(new_key, val, time.Hour*24) /* SET overwrites the value if it exists */
	return true
}

func (myredis_client *RedisDb) Get(stream string, key string) (bool, string) {
	new_key := key + stream
	value := myredis_client.R.Get(new_key)
	if value == nil || value.Val() == "" {
		fmt.Println("FASTMEM_GET: failed", new_key)
		return false, ""
	}

	return true, value.Val()
}

func (myredis_client *RedisDb) Del(stream string, key string) (bool, string) {
	new_key := key + stream
	value := myredis_client.R.Get(new_key)
	if value == nil || value.Val() == "" {
		return false, ""
	}
	myredis_client.R.Del(new_key)
	return true, value.Val()
}

/* Advise : Do not use this func unless you know what you are doing */
func (myredis_client *RedisDb) BlockedPush(stream, key, value string) {
	if v := myredis_client.R.RPush(stream, key, value); v.Err() != nil {
		fmt.Println("failed to push, error:", v.String())
	}
}

/* Advise : Do not use this func unless you know what you are doing */
func (myredis_client *RedisDb) BlockedPop(stream string) (bool, []string) {
	v := myredis_client.R.BLPop(0, stream)
	if v.Err() != nil {
		return false, v.Val()
	}
	return true, v.Val()
}

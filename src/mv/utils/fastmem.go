package utils

import (
	"fmt"
	"time"

	redis "gopkg.in/redis.v3"
)

// RedisDb is a StorageClass for redis handler
type RedisDb struct {
	R *redis.Client
}

// FastMemInit initializes the redis-handler.
func FastMemInit(redisIP string) (bool, *RedisDb) {
	dial := redisIP + ":6379"
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

// FastMemDestroy destroys the fastmem handler
func (obj *RedisDb) FastMemDestroy() {
	obj.R.Close()
}

// TimedAdd register entries into fastmem for 2 days
func (obj *RedisDb) TimedAdd(stream string, key string, val string) bool {
	addKey := key + stream
	obj.R.Set(addKey, val, time.Hour*24) /* SET overwrites the value if it exists */
	return true
}

// Get returns entries against the key from fastmem. It may return false.
func (obj *RedisDb) Get(stream string, key string) (bool, string) {
	getKey := key + stream
	value := obj.R.Get(getKey)
	if value == nil || value.Val() == "" {
		fmt.Println("FASTMEM_GET: failed", getKey)
		return false, ""
	}

	return true, value.Val()
}

// Del deletes the entry from fastmem. If found.
func (obj *RedisDb) Del(stream string, key string) (bool, string) {
	delKey := key + stream
	value := obj.R.Get(delKey)
	if value == nil || value.Val() == "" {
		return false, ""
	}
	obj.R.Del(delKey)
	return true, value.Val()
}

// BlockedPush Advise : Do not use this func unless you know what you are doing */
func (obj *RedisDb) BlockedPush(stream, key, value string) {
	if v := obj.R.RPush(stream, key, value); v.Err() != nil {
		fmt.Println("failed to push, error:", v.String())
	}
}

// BlockedPop Advise : Do not use this func unless you know what you are doing */
func (obj *RedisDb) BlockedPop(stream string) (bool, []string) {
	v := obj.R.BLPop(0, stream)
	if v.Err() != nil {
		return false, v.Val()
	}
	return true, v.Val()
}

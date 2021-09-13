package cache

import (
	"context"
	json2 "encoding/json"
	"github.com/go-redis/redis/v8"
	"redisTesting/models"
	"time"
)

type redisCache struct {
	host string
	db int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) PostCache {
	return &redisCache{
		host: host,
		db: db,
		expires: expires,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		Password: "",
		DB: cache.db,
	})
}

func (cache *redisCache) Set(key string, value *models.Post){
	client := cache.getClient()
	json, err := json2.Marshal(value)
	if err != nil{
		panic(err.Error())
	}

	client.Set(context.Background(), key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *models.Post{
	client := cache.getClient()
	val, err := client.Get(context.Background(), key).Result()
	if err != nil{
		panic(err.Error())
		return nil
	}
	var post models.Post
	err = json2.Unmarshal([]byte(val), &post)
	if err !=nil{
		panic(err.Error())
		return nil
	}
	return &post
}

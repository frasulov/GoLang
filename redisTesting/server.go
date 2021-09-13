package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redisTesting/cache"
	"redisTesting/models"
)

func main() {

	server := gin.Default()

	redis := cache.NewRedisCache("localhost:6379", 1, 1000)
	fmt.Println("connceted ro redis")
	server.POST("/posts", func(c *gin.Context) {
		var post models.Post
		c.ShouldBindJSON(&post)
		fmt.Println("here: ", post)
		redis.Set("post-1", &post)
		fmt.Println("post setted!")
	})

	server.GET("/posts", func(context *gin.Context) {
		post := redis.Get("post-1")
		fmt.Println("Post fetched")
		fmt.Println(post)
	})
	server.Run(":9001")
}

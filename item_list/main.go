package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v3"
)

func main() {
	client := redis.NewClient(&redis.Options{
		//Addr:     "localhost:6379",
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	router := gin.Default()
	router.Run(":3005")
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v7"
)

func main() {
	// connecting to redis database
	fmt.Println("Connecting to local redis database")
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	fmt.Println("Starting go-compose")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from go-compose!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "pong",
		})
	})

	// starting server
	r.Run()
}

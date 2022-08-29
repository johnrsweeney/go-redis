package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func getStatus(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379",
		Password: "",
		DB: 0,
	})

	status, err := client.Get("status").Result()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H {
		"status": status,
	})
}

func setStatus(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379",
		Password: "",
		DB: 0,
	})

	status_param := c.Request.URL.Query().Get("status");

	err := client.Set("status", status_param, 0).Err();
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	router := gin.Default()
//	router.SetTrustedProxies([]string{"172.0.0.1"})

	router.GET("/setstatus", setStatus)
	router.GET("/status", getStatus)

	router.Run(":80")
}


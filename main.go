package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type album struct {
	ID		string	`json:"id"`
	Title	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

var albums = []album {
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getStatus(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379",
		Password: "",
		DB: 0,
	})

	name, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H {
		"status": name,
	})
}

func main() {
	router := gin.Default()
//	router.SetTrustedProxies([]string{"172.0.0.1"})

	router.GET("/status", getStatus)

	router.Run(":80")
}


package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
	docs "github.com/johnrsweeney/go-redis/docs"
)

// @title Status API
// @version 0.1.0
// @description Cruising' and Bruisin'

// @host 18.170.214.52
// @BaseBath /

// @Summary get status
// @Schemeds
// @Descriptions no dunking
// @Tags noTag
// @Accept json
// @Produce json
// Success 200
// @Router /status [get]
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
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/setstatus", setStatus)
	router.GET("/status", getStatus)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("0.0.0.0:8080")
}


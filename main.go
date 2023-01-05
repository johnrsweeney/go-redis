package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	//docs "github.com/johnrsweeney/go-redis/docs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	//swaggerfiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title       Status API
// @version     0.1.0
// @description Cruising' and Bruisin'

// @host     18.170.214.52
// @BaseBath /

// @Summary get status
// @Schemes
// @Description no dunking
// @Tags        noTag
// @Accept      json
// @Produce     json
// Success 200
// @Router      /status [get]
func getStatus(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})

	status, err := client.Get("status").Result()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})

	requestsProcessed.Inc()
}

// @Summary set status
// @Schemes
// @Description no bare feet
// @tags        noTag
// @Accept      json
// @Param       status	query string false "Value to set status"
// @Produce     json
// @Success     200
// @Router      /setstatus [get]
func setStatus(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})

	status_param := c.Request.URL.Query().Get("status")

	err := client.Set("status", status_param, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}

var (
	requestsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "requests_processed",
		Help: "The total of number status requests processed",
	})
)

func main() {
	router := gin.Default()
	//docs.SwaggerInfo.BasePath = "/"

	//router.GET("/", func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/swagger/*any")
	//})

	router.GET("/setstatus", setStatus)
	router.GET("/status", getStatus)

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Run("0.0.0.0:80")
}

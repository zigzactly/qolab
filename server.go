package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := Router()
	StartServer(r)
}

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "home",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func StartServer(r *gin.Engine) {
	r.Run(":8080")
}

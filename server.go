package main

import "github.com/gin-gonic/gin"

func GinNewServer() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
func RunServer() error {
	r := GinNewServer()
	return r.Run()
}

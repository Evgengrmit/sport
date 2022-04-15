package main

import "github.com/gin-gonic/gin"

func GinNewServer() *gin.Engine {
	router := gin.Default()
	return router
}

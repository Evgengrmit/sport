package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"sport/pkg/service"
)

type Handler struct {
	services *service.Service
}
type errorResponse struct {
	Message string `json:"message"`
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

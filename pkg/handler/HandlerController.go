package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	schedules := router.Group("/schedules")
	{
		schedules.GET("/", h.GetAllComplexes)
		schedules.GET("/1/", h.GetAllSchedules)
	}
	return router
}

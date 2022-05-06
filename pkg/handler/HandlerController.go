package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	wods := router.Group("/wod")
	{
		wods.GET("/", h.GetAllWorkoutDays)
		wods.GET("/latest", h.GetWorkoutLatest)

	}
	schedules := router.Group("/schedules")
	{
		schedules.GET("/", h.GetAllSchedules)
	}
	return router
}

package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/login", h.Login)
	wods := router.Group("/wod")
	{
		wods.GET("/", h.GetAllWorkoutDays)
		wods.GET("/latest", h.GetWorkoutLatest)

	}
	schedules := router.Group("/schedules")
	{
		schedules.GET("/", h.GetAllSchedules)
	}
	exercise := router.Group("/exercise")
	{
		exercise.POST("/result", h.AddResult)
	}
	exercises := router.Group("/exercises")
	{
		exercises.GET("/", h.GetAllExercises)
	}
	return router
}

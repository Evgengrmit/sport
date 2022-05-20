package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	login := router.Group("/login")
	{
		login.POST("/", h.Login)
		login.POST("/verify", h.Verify)
	}
	wods := router.Group("/wod")
	{
		wods.GET("/", h.GetAllWorkoutDays)
		wods.GET("/latest", h.GetWorkoutLatest)

	}
	workout := router.Group("/workout")
	{
		workout.POST("/result", h.AddWorkoutResult)
	}
	schedules := router.Group("/schedules")
	{
		schedules.GET("/", h.GetAllSchedules)
	}
	exercise := router.Group("/exercise")
	{
		exercise.POST("/result", h.AddExerciseResult)
	}
	exercises := router.Group("/exercises")
	{
		exercises.GET("/", h.GetAllExercises)
	}
	return router
}

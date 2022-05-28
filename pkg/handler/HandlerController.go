package handler

import (
	"github.com/gin-gonic/gin"
	"sport/pkg/repository/schedulesRepo"
)

func (h *Handler) InitRoutes() *gin.Engine {

	storageRootPath := schedulesRepo.GetStorageRootPath()

	router := gin.Default()

	router.Static(schedulesRepo.GetStorageRootUrl(), storageRootPath)

	router.POST("/login", h.Login)
	wods := router.Group("/wod")
	{
		wods.GET("/", h.GetAllWorkoutDays)
		wods.GET("/latest", h.GetWorkoutLatest)

	}
	workout := router.Group("/workout")
	{
		workout.GET("/:id/result", h.GetWorkoutResults)
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

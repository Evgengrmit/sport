package handler

import (
	"github.com/gin-gonic/gin"
	imageUtils "sport/pkg/utils"
)

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.Default()

	storageRootPath := imageUtils.GetStorageRootPath()
	router.Static(imageUtils.GetStorageUrlPrefix(), storageRootPath)

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

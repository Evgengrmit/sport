package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sport/pkg/repository/authRepo"
	"sport/pkg/repository/workoutResultRepo"
	"sport/pkg/service"
)

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
func (h *Handler) Login(c *gin.Context) {
	var input authRepo.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": result})
}
func (h *Handler) GetAllWorkoutDays(c *gin.Context) {
	complexes, err := h.services.WorkoutDay.GetAllWorkoutDays()

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllWorkoutDaysResponse{Data: complexes})
}
func (h *Handler) GetWorkoutLatest(c *gin.Context) {
	workoutDaysByDay, err := h.services.WorkoutDay.GetWorkoutLatest()
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": workoutDaysByDay})
}

func (h *Handler) GetAllSchedules(c *gin.Context) {
	schedules, err := h.services.Schedule.GetAllSchedules()

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllSchedulesResponse{Data: schedules})
}

func (h *Handler) GetAllExercises(c *gin.Context) {
	exercises, err := h.services.Exercise.GetAllExercises()
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": exercises})
}
func (h *Handler) AddWorkoutResult(c *gin.Context) {
	var input workoutResultRepo.WorkoutResult
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.services.WorkoutResult.CreateWorkoutResult(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "ok"})
}

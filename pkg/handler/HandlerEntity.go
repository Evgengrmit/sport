package handler

import (
	"sport/pkg/repository/schedulesRepo"
	"sport/pkg/repository/wodRepo"
	"sport/pkg/service"
)

type Handler struct {
	services *service.Service
}

type statusResponse struct {
	Status string `json:"status"`
}
type errorResponse struct {
	Message string `json:"message"`
}
type getAllWorkoutDaysResponse struct {
	Data []wodRepo.WorkoutDay `json:"data"`
}
type getAllSchedulesResponse struct {
	Data []schedulesRepo.Schedule `json:"data"`
}

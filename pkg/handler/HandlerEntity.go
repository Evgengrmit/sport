package handler

import (
	"sport/pkg/service"
	"sport/sportclub/schedules"
	"sport/sportclub/wod"
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
	Data []wod.WorkoutDayJSON `json:"data"`
}
type getAllSchedulesResponse struct {
	Data []schedules.ScheduleJSON `json:"data"`
}

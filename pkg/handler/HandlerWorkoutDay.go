package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllWorkoutDays(c *gin.Context) {
	complexes, err := h.services.WorkoutDay.GetAllWorkoutDays()

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllComplexResponse{Data: complexes})
}

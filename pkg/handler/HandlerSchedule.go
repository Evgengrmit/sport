package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllSchedules(c *gin.Context) {
	schedules, err := h.services.Schedule.GetAllSchedules()

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, schedules)
}

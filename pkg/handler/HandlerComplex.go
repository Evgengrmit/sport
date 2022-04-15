package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllComplexes(c *gin.Context) {
	fmt.Println("dfdfd")
	complexes, err := h.services.Complex.GetAllComplexes()

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllComplexResponse{Data: complexes})
}

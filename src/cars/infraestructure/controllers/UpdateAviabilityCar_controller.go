package controllers

import (
	"ejercicio1/src/cars/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateAvailabilityCarController struct {
	usecase application.UpdateAvailabilityCar
}

func NewUpdateAvailabilityCarController(usecase application.UpdateAvailabilityCar) *UpdateAvailabilityCarController {
	return &UpdateAvailabilityCarController{usecase: usecase}
}

func (uc_c *UpdateAvailabilityCarController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request struct {
		Available bool `json:"available"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = uc_c.usecase.Execute(id, request.Available)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car availability updated successfully"})
}

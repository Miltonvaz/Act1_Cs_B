package controllers

import (
	"ejercicio1/src/Appointment/application"
	"ejercicio1/src/Appointment/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EditAppointmentController struct {
	usecase application.EditAppointment
}

func NewEditAppointmentController(usecase application.EditAppointment) *EditAppointmentController {
	return &EditAppointmentController{usecase: usecase}
}

func (eac *EditAppointmentController) Execute(c *gin.Context) {
	var appointment entities.TestDriveAppointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := eac.usecase.Execute(appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment updated successfully"})
}

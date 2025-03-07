package controllers

import (
	"ejercicio1/src/Appointment/application"
	"ejercicio1/src/Appointment/domain"
	"ejercicio1/src/Appointment/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateAppointmentController struct {
	usecase         application.CreateAppointment
	appointmentRepo domain.IAppointment
}

func NewCreateAppointmentController(usecase application.CreateAppointment, appointmentRepo domain.IAppointment) *CreateAppointmentController {
	return &CreateAppointmentController{
		usecase:         usecase,
		appointmentRepo: appointmentRepo,
	}
}

func (cac *CreateAppointmentController) Execute(c *gin.Context) {
	var appointment entities.TestDriveAppointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if !appointment.TestDate.Time.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Test date must be in the future"})
		return
	}

	validStatuses := []string{"pending", "confirmed", "canceled"}
	statusValid := false
	for _, validStatus := range validStatuses {
		if appointment.Status == validStatus {
			statusValid = true
			break
		}
	}

	if !statusValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	createdAppointment, err := cac.usecase.Execute(appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdAppointment)
}

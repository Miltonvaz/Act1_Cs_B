package controllers

import (
	"ejercicio1/src/Appointment/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateAppointmentStatusController struct {
	usecase application.UpdateAppointmentStatus
}

func NewUpdateAppointmentStatusController(usecase application.UpdateAppointmentStatus) *UpdateAppointmentStatusController {
	return &UpdateAppointmentStatusController{usecase: usecase}
}

func (uasc *UpdateAppointmentStatusController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var requestBody struct {
		Status string `json:"status"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	status := requestBody.Status
	if err := uasc.usecase.Execute(id, status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment status updated successfully"})
}

func (uasc *UpdateAppointmentStatusController) GetStatus(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	appointment, err := uasc.usecase.GetStatus(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": appointment.Status})
}

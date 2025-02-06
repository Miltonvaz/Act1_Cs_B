package controllers

import (
	"ejercicio1/src/Appointment/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAppointmentStatusController struct {
	usecase application.ViewAppointmentStatus
}

func NewGetAppointmentStatusController(usecase application.ViewAppointmentStatus) *GetAppointmentStatusController {
	return &GetAppointmentStatusController{usecase: usecase}
}

func (gasc *GetAppointmentStatusController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	status, err := gasc.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}

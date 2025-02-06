package controllers

import (
	"ejercicio1/src/Appointment/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteAppointmentController struct {
	usecase application.DeleteAppointment
}

func NewDeleteAppointmentController(usecase application.DeleteAppointment) *DeleteAppointmentController {
	return &DeleteAppointmentController{usecase: usecase}
}

func (dac *DeleteAppointmentController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := dac.usecase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted successfully"})
}

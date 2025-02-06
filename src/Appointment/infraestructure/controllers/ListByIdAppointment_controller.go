package controllers

import (
	"ejercicio1/src/Appointment/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewAppointmentByIdController struct {
	usecase application.ViewAppointmentById
}

func NewViewAppointmentByIdController(usecase application.ViewAppointmentById) *ViewAppointmentByIdController {
	return &ViewAppointmentByIdController{usecase: usecase}
}

func (vac *ViewAppointmentByIdController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	appointment, err := vac.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

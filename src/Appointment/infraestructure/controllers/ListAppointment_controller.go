package controllers

import (
	"ejercicio1/src/Appointment/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListAppointmentsController struct {
	usecase application.ListAppointments
}

func NewListAppointmentsController(usecase application.ListAppointments) *ListAppointmentsController {
	return &ListAppointmentsController{usecase: usecase}
}

func (lac *ListAppointmentsController) Execute(c *gin.Context) {
	appointments, err := lac.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

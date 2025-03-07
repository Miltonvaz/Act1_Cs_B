package routes_a

import (
	"ejercicio1/src/Appointment/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAppointmentRoutes(

	r *gin.Engine,
	createAppointmentController *controllers.CreateAppointmentController,
	listAppointmentController *controllers.ListAppointmentsController,
	viewByIdAppointmentController *controllers.ViewAppointmentByIdController,
	updateAppointmentController *controllers.EditAppointmentController,
	deleteAppointmentController *controllers.DeleteAppointmentController,
	updateAppointmentStatusController *controllers.UpdateAppointmentStatusController,
	getAppointmentStatusController *controllers.GetAppointmentStatusController,

) {
	r.POST("/appointments", createAppointmentController.Execute)
	r.GET("/appointments", listAppointmentController.Execute)
	r.GET("/appointments/:id", viewByIdAppointmentController.Execute)
	r.PUT("/appointments/:id", updateAppointmentController.Execute)
	r.DELETE("/appointments/:id", deleteAppointmentController.Execute)
	r.PUT("/appointments/:id/status", updateAppointmentStatusController.Execute)
	r.GET("/appointments/:id/status", getAppointmentStatusController.Execute)

}

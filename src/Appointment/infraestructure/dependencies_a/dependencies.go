package dependencies_a

import (
	"ejercicio1/src/Appointment/application"
	"ejercicio1/src/Appointment/infraestructure/adapter"
	"ejercicio1/src/Appointment/infraestructure/controllers"
	"ejercicio1/src/Appointment/infraestructure/db"
	"ejercicio1/src/core"
)

func InitAppointments() (
	*controllers.CreateAppointmentController,
	*controllers.ListAppointmentsController,
	*controllers.DeleteAppointmentController,
	*controllers.EditAppointmentController,
	*controllers.ViewAppointmentByIdController,
	*controllers.UpdateAppointmentStatusController,
	*controllers.GetAppointmentStatusController,
	*controllers.NotificationController,
	error,
) {

	pool := core.GetDBPool()
	ps := db.NewMySQL(pool.DB)

	notificationAdapter := adapter.NewNotificationAdapter()

	createAppointment := application.NewCreateAppointment(ps, notificationAdapter)
	listAppointment := application.NewListAppointments(ps)
	editAppointment := application.NewEditAppointment(ps)
	deleteAppointment := application.NewDeleteAppointment(ps)
	viewByIdAppointment := application.NewViewAppointmentById(ps)
	updateAppointmentStatus := application.NewUpdateAppointmentStatus(ps)
	viewAppointmentStatus := application.NewViewAppointmentStatus(ps)

	createAppointmentController := controllers.NewCreateAppointmentController(*createAppointment)
	listAppointmentController := controllers.NewListAppointmentsController(*listAppointment)
	editAppointmentController := controllers.NewEditAppointmentController(*editAppointment)
	deleteAppointmentController := controllers.NewDeleteAppointmentController(*deleteAppointment)
	viewByIdAppointmentController := controllers.NewViewAppointmentByIdController(*viewByIdAppointment)
	updateAppointmentStatusController := controllers.NewUpdateAppointmentStatusController(*updateAppointmentStatus)
	viewAppointmentStatusController := controllers.NewGetAppointmentStatusController(*viewAppointmentStatus)

	notificationController := controllers.NewNotificationController(notificationAdapter)

	return createAppointmentController, listAppointmentController, deleteAppointmentController,
		editAppointmentController, viewByIdAppointmentController, updateAppointmentStatusController,
		viewAppointmentStatusController, notificationController, nil
}

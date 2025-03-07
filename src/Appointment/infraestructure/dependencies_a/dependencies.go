package dependencies_a

import (
	"ejercicio1/src/Appointment/application"
	"ejercicio1/src/Appointment/application/repositories"
	"ejercicio1/src/Appointment/infraestructure/adapters"
	"ejercicio1/src/Appointment/infraestructure/adapters/a_rabbit"
	"ejercicio1/src/Appointment/infraestructure/controllers"
	"ejercicio1/src/core"
	"log"
)

func InitAppointments() (
	*controllers.CreateAppointmentController,
	*controllers.ListAppointmentsController,
	*controllers.DeleteAppointmentController,
	*controllers.EditAppointmentController,
	*controllers.ViewAppointmentByIdController,
	*controllers.UpdateAppointmentStatusController,
	*controllers.GetAppointmentStatusController,
	error,
) {

	pool := core.GetDBPool()
	ps := adapters.NewMySQL(pool.DB)

	rabbitMQAdapter, err := a_rabbit.NewRabbitMQAdapter()
	if err != nil {
		log.Printf("Error inicializando RabbitMQ: %v", err)
		return nil, nil, nil, nil, nil, nil, nil, err
	}

	serviceNotification := repositories.NewServiceNotification(rabbitMQAdapter)

	createAppointment := application.NewCreateAppointment(ps, serviceNotification)
	listAppointment := application.NewListAppointments(ps)
	editAppointment := application.NewEditAppointment(ps)
	deleteAppointment := application.NewDeleteAppointment(ps)
	viewByIdAppointment := application.NewViewAppointmentById(ps)
	updateAppointmentStatus := application.NewUpdateAppointmentStatus(ps)
	viewAppointmentStatus := application.NewViewAppointmentStatus(ps)

	createAppointmentController := controllers.NewCreateAppointmentController(*createAppointment, ps)
	listAppointmentController := controllers.NewListAppointmentsController(*listAppointment)
	editAppointmentController := controllers.NewEditAppointmentController(*editAppointment)
	deleteAppointmentController := controllers.NewDeleteAppointmentController(*deleteAppointment)
	viewByIdAppointmentController := controllers.NewViewAppointmentByIdController(*viewByIdAppointment)
	updateAppointmentStatusController := controllers.NewUpdateAppointmentStatusController(*updateAppointmentStatus)
	viewAppointmentStatusController := controllers.NewGetAppointmentStatusController(*viewAppointmentStatus)

	return createAppointmentController, listAppointmentController, deleteAppointmentController,
		editAppointmentController, viewByIdAppointmentController,
		updateAppointmentStatusController, viewAppointmentStatusController, nil
}

package application

import (
	"ejercicio1/src/Appointment/application/repositories"
	"ejercicio1/src/Appointment/domain"
	"ejercicio1/src/Appointment/domain/entities"
	"log"
)

type CreateAppointment struct {
	appointmentRepo     domain.IAppointment
	serviceNotification *repositories.ServiceNotification
}

func NewCreateAppointment(appointmentRepo domain.IAppointment, serviceNotification *repositories.ServiceNotification) *CreateAppointment {
	return &CreateAppointment{
		appointmentRepo:     appointmentRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *CreateAppointment) Execute(appointment entities.TestDriveAppointment) (entities.TestDriveAppointment, error) {

	created, err := c.appointmentRepo.Save(appointment)

	err = c.serviceNotification.NotifyAppointmentCreated(created)
	if err != nil {
		log.Printf("Error notificando cita creada: %v", err)
		return entities.TestDriveAppointment{}, err
	}

	return created, nil
}

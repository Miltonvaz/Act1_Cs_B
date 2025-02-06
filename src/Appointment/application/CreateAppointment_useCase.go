package application

import (
	"ejercicio1/src/Appointment/domain"
	"ejercicio1/src/Appointment/domain/entities"
	"ejercicio1/src/Appointment/infraestructure/adapter"
)

type CreateAppointment struct {
	appointmentRepo     domain.IAppointment
	notificationAdapter *adapter.NotificationAdapter
}

func NewCreateAppointment(appointmentRepo domain.IAppointment, notificationAdapter *adapter.NotificationAdapter) *CreateAppointment {
	return &CreateAppointment{
		appointmentRepo:     appointmentRepo,
		notificationAdapter: notificationAdapter,
	}
}

func (c *CreateAppointment) Execute(appointment entities.TestDriveAppointment) (entities.TestDriveAppointment, error) {
	created, err := c.appointmentRepo.Save(appointment)
	if err != nil {
		return entities.TestDriveAppointment{}, err
	}

	c.notificationAdapter.Notify(created)

	return created, nil
}

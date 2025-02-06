package application

import (
	"ejercicio1/src/Appointment/domain"
	"ejercicio1/src/Appointment/domain/entities"
)

type EditAppointment struct {
	appointmentRepo domain.IAppointment
}

func NewEditAppointment(appointmentRepo domain.IAppointment) *EditAppointment {
	return &EditAppointment{appointmentRepo: appointmentRepo}
}

func (ea *EditAppointment) Execute(appointment entities.TestDriveAppointment) error {
	return ea.appointmentRepo.Edit(appointment)
}

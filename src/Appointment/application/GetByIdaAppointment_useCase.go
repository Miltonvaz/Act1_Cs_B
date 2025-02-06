package application

import (
	"ejercicio1/src/Appointment/domain"
	"ejercicio1/src/Appointment/domain/entities"
)

type ViewAppointmentById struct {
	appointmentRepo domain.IAppointment
}

func NewViewAppointmentById(appointmentRepo domain.IAppointment) *ViewAppointmentById {
	return &ViewAppointmentById{appointmentRepo: appointmentRepo}
}

func (va *ViewAppointmentById) Execute(id int) (entities.TestDriveAppointment, error) {
	return va.appointmentRepo.GetById(id)
}

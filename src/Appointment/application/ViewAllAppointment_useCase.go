package application

import (
	"ejercicio1/src/Appointment/domain"
	"ejercicio1/src/Appointment/domain/entities"
)

type ListAppointments struct {
	appointmentRepo domain.IAppointment
}

func NewListAppointments(appointmentRepo domain.IAppointment) *ListAppointments {
	return &ListAppointments{appointmentRepo: appointmentRepo}
}

func (la *ListAppointments) Execute() ([]entities.TestDriveAppointment, error) {
	return la.appointmentRepo.GetAll()
}

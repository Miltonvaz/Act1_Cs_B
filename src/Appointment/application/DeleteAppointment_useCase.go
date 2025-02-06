package application

import (
	"ejercicio1/src/Appointment/domain"
)

type DeleteAppointment struct {
	appointmentRepo domain.IAppointment
}

func NewDeleteAppointment(appointmentRepo domain.IAppointment) *DeleteAppointment {
	return &DeleteAppointment{appointmentRepo: appointmentRepo}
}

func (da *DeleteAppointment) Execute(id int) error {
	return da.appointmentRepo.Delete(id)
}

package application

import (
	"ejercicio1/src/Appointment/domain"
	"ejercicio1/src/Appointment/domain/entities"
)

type UpdateAppointmentStatus struct {
	appointmentRepo domain.IAppointment
}

func NewUpdateAppointmentStatus(appointmentRepo domain.IAppointment) *UpdateAppointmentStatus {
	return &UpdateAppointmentStatus{appointmentRepo: appointmentRepo}
}

func (ua *UpdateAppointmentStatus) Execute(id int, status string) error {

	return ua.appointmentRepo.UpdateStatus(id, status)
}

func (ua *UpdateAppointmentStatus) GetStatus(id int) (*entities.TestDriveAppointment, error) {

	appointment, err := ua.appointmentRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

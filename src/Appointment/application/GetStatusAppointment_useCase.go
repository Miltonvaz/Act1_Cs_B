package application

import "ejercicio1/src/Appointment/domain"

type ViewAppointmentStatus struct {
	appointmentRepo domain.IAppointment
}

func NewViewAppointmentStatus(appointmentRepo domain.IAppointment) *ViewAppointmentStatus {
	return &ViewAppointmentStatus{appointmentRepo: appointmentRepo}
}

func (vas *ViewAppointmentStatus) Execute(id int) (string, error) {
	status, err := vas.appointmentRepo.GetStatus(id)
	if err != nil {
		return "", err
	}
	return status, nil
}

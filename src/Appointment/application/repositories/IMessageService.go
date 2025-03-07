package repositories

import "ejercicio1/src/Appointment/domain/entities"

type NotificationPort interface {
	PublishEvent(eventType string, appointment entities.TestDriveAppointment) error
}

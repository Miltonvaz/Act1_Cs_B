package domain

import "ejercicio1/src/Appointment/domain/entities"

type IAppointment interface {
	Save(entities.TestDriveAppointment) (entities.TestDriveAppointment, error)
	GetAll() ([]entities.TestDriveAppointment, error)
	GetById(id int) (entities.TestDriveAppointment, error)
	Edit(entities.TestDriveAppointment) error
	UpdateStatus(id int, status string) error
	GetStatus(id int) (string, error)
	Delete(id int) error
}

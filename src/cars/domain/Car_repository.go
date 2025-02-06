package domain

import "ejercicio1/src/cars/domain/entities"

type ICar interface {
	Save(entities.Car) (entities.Car, error)
	GetAll() ([]entities.Car, error)
	GetById(id int) (entities.Car, error)
	Edit(entities.Car) error
	UpdateAvailability(id int, available bool) error
	GetAvailable() ([]entities.Car, error)
	Delete(id int) error
}

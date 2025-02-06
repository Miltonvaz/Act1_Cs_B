package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
	"ejercicio1/src/cars/infraestructure/db"
)

type CreateCar struct {
	db domain.ICar
}

func NewCreateCar(db *db.MySQL) *CreateCar {
	return &CreateCar{db: db}
}

func (cc *CreateCar) Execute(car entities.Car) (entities.Car, error) {
	return cc.db.Save(car)
}

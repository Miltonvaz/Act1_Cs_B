package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
	"ejercicio1/src/cars/infraestructure/db"
)

type ListCar struct {
	db domain.ICar
}

func NewListCar(db *db.MySQL) *ListCar {
	return &ListCar{db: db}
}

func (lc ListCar) Execute() ([]entities.Car, error) {
	return lc.db.GetAll()
}

package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
	"ejercicio1/src/cars/infraestructure/db"
)

type ViewByIdCar struct {
	db domain.ICar
}

func NewViewByIdCar(db *db.MySQL) *ViewByIdCar {
	return &ViewByIdCar{db: db}
}

func (vc ViewByIdCar) Execute(id int) (entities.Car, error) {
	return vc.db.GetById(id)

}

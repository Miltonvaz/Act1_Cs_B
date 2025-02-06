package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/infraestructure/db"
)

type DeleteCar struct {
	db domain.ICar
}

func NewDeleteCar(db *db.MySQL) *DeleteCar {
	return &DeleteCar{db: db}
}

func (dc DeleteCar) Execute(id int) error {
	return dc.db.Delete(id)
}

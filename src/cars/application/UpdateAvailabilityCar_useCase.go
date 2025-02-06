package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/infraestructure/db"
)

type UpdateAvailabilityCar struct {
	db domain.ICar
}

func NewUpdateAvailabilityCar(db *db.MySQL) *UpdateAvailabilityCar {
	return &UpdateAvailabilityCar{db: db}
}

func (uc *UpdateAvailabilityCar) Execute(id int, available bool) error {
	return uc.db.UpdateAvailability(id, available)
}

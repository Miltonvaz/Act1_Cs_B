package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
	"fmt"
)

type GetAvailableCars struct {
	repository domain.ICar
}

func NewGetAvailableCars(repository domain.ICar) *GetAvailableCars {
	return &GetAvailableCars{repository: repository}
}

func (uc *GetAvailableCars) Execute() ([]entities.Car, error) {

	cars, err := uc.repository.GetAvailable()
	if err != nil {
		return nil, fmt.Errorf("failed to get available cars: %v", err)
	}

	return cars, nil
}

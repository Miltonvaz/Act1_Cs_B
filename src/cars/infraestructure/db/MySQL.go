package db

import (
	"database/sql"
	"ejercicio1/src/cars/domain/entities"
	"fmt"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) GetAvailable() ([]entities.Car, error) {
	query := "SELECT id, make, model, year, mileage, fuel_type, available FROM cars WHERE available = 1"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve available cars: %v", err)
	}
	defer rows.Close()

	var cars []entities.Car
	for rows.Next() {
		var car entities.Car
		err := rows.Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.Mileage, &car.FuelType, &car.Available)
		if err != nil {
			return nil, fmt.Errorf("failed to scan car row: %v", err)
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (m *MySQL) Save(car entities.Car) (entities.Car, error) {
	query := "INSERT INTO cars (make, model, year, mileage, fuel_type, available) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := m.conn.Exec(query, car.Make, car.Model, car.Year, car.Mileage, car.FuelType, car.Available)
	if err != nil {
		return entities.Car{}, fmt.Errorf("failed to save car: %v", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return entities.Car{}, fmt.Errorf("failed to get last insert ID: %v", err)
	}
	car.ID = int32(lastInsertID)

	return car, nil
}

func (m *MySQL) GetAll() ([]entities.Car, error) {
	query := "SELECT id, make, model, year, mileage, fuel_type, available FROM cars"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cars: %v", err)
	}
	defer rows.Close()

	var cars []entities.Car
	for rows.Next() {
		var car entities.Car
		err := rows.Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.Mileage, &car.FuelType, &car.Available)
		if err != nil {
			return nil, fmt.Errorf("failed to scan car row: %v", err)
		}
		cars = append(cars, car)
	}

	return cars, nil
}

func (m *MySQL) GetById(id int) (entities.Car, error) {
	query := "SELECT id, make, model, year, mileage, fuel_type, available FROM cars WHERE id = ?"
	row := m.conn.QueryRow(query, id)

	var car entities.Car
	err := row.Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.Mileage, &car.FuelType, &car.Available)
	if err == sql.ErrNoRows {
		return entities.Car{}, fmt.Errorf("car not found")
	} else if err != nil {
		return entities.Car{}, fmt.Errorf("failed to retrieve car: %v", err)
	}

	return car, nil
}

func (m *MySQL) Edit(car entities.Car) error {
	query := "UPDATE cars SET make = ?, model = ?, year = ?, mileage = ?, fuel_type = ?, available = ? WHERE id = ?"
	_, err := m.conn.Exec(query, car.Make, car.Model, car.Year, car.Mileage, car.FuelType, car.Available, car.ID)
	if err != nil {
		return fmt.Errorf("failed to update car: %v", err)
	}
	return nil
}

func (m *MySQL) UpdateAvailability(id int, available bool) error {
	query := "UPDATE cars SET available = ? WHERE id = ?"
	_, err := m.conn.Exec(query, available, id)
	if err != nil {
		return fmt.Errorf("failed to update car availability: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM cars WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete car: %v", err)
	}
	return nil
}

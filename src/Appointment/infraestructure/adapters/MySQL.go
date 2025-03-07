package adapters

import (
	"database/sql"
	"ejercicio1/src/Appointment/domain/entities"
	"fmt"
	"time"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(appointment entities.TestDriveAppointment) (entities.TestDriveAppointment, error) {
	query := "INSERT INTO appointments (client_id, car_id, test_date, location, status) VALUES (?, ?, ?, ?, ?)"
	result, err := m.conn.Exec(query, appointment.ClientID, appointment.CarID, appointment.TestDate, appointment.Location, appointment.Status)
	if err != nil {
		return entities.TestDriveAppointment{}, fmt.Errorf("failed to save appointment: %v", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return entities.TestDriveAppointment{}, fmt.Errorf("failed to get last insert ID: %v", err)
	}
	appointment.AppointmentID = int(lastInsertID)

	return appointment, nil
}

func (m *MySQL) GetAll() ([]entities.TestDriveAppointment, error) {
	query := "SELECT appointment_id, client_id, car_id, test_date, location, status FROM appointments"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve appointments: %v", err)
	}
	defer rows.Close()

	var appointments []entities.TestDriveAppointment
	for rows.Next() {
		var appointment entities.TestDriveAppointment
		var testDate []byte
		err := rows.Scan(&appointment.AppointmentID, &appointment.ClientID, &appointment.CarID, &testDate, &appointment.Location, &appointment.Status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan appointment row: %v", err)
		}

		if len(testDate) > 0 {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", string(testDate))
			if err != nil {
				return nil, fmt.Errorf("failed to parse test date: %v", err)
			}
			appointment.TestDate = sql.NullTime{Time: parsedTime, Valid: true}
		} else {
			appointment.TestDate = sql.NullTime{Valid: false}
		}

		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (m *MySQL) GetStatus(id int) (string, error) {
	query := "SELECT status FROM appointments WHERE appointment_id = ?"
	row := m.conn.QueryRow(query, id)

	var status string
	err := row.Scan(&status)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("appointment not found")
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve appointment status: %v", err)
	}

	return status, nil
}

func (m *MySQL) GetById(id int) (entities.TestDriveAppointment, error) {
	query := "SELECT appointment_id, client_id, car_id, test_date, location, status FROM appointments WHERE appointment_id = ?"
	row := m.conn.QueryRow(query, id)

	var appointment entities.TestDriveAppointment
	var testDate []byte
	err := row.Scan(&appointment.AppointmentID, &appointment.ClientID, &appointment.CarID, &testDate, &appointment.Location, &appointment.Status)
	if err == sql.ErrNoRows {
		return entities.TestDriveAppointment{}, fmt.Errorf("appointment not found")
	} else if err != nil {
		return entities.TestDriveAppointment{}, fmt.Errorf("failed to retrieve appointment: %v", err)
	}

	if len(testDate) > 0 {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(testDate))
		if err != nil {
			return entities.TestDriveAppointment{}, fmt.Errorf("failed to parse test date: %v", err)
		}
		appointment.TestDate = sql.NullTime{Time: parsedTime, Valid: true}
	} else {
		appointment.TestDate = sql.NullTime{Valid: false}
	}

	return appointment, nil
}

func (m *MySQL) Edit(appointment entities.TestDriveAppointment) error {
	query := "UPDATE appointments SET client_id = ?, car_id = ?, test_date = ?, location = ?, status = ? WHERE appointment_id = ?"
	_, err := m.conn.Exec(query, appointment.ClientID, appointment.CarID, appointment.TestDate, appointment.Location, appointment.Status, appointment.AppointmentID)
	if err != nil {
		return fmt.Errorf("failed to update appointment: %v", err)
	}
	return nil
}

func (m *MySQL) UpdateStatus(id int, status string) error {
	query := "UPDATE appointments SET status = ? WHERE appointment_id = ?"
	_, err := m.conn.Exec(query, status, id)
	if err != nil {
		return fmt.Errorf("failed to update appointment status: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM appointments WHERE appointment_id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete appointment: %v", err)
	}
	return nil
}

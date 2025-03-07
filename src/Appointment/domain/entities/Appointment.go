package entities

import (
	"database/sql"
)

type TestDriveAppointment struct {
	AppointmentID int          `json:"appointment_id"`
	CarID         int          `json:"car_id"`
	ClientID      int          `json:"client_id"`
	TestDate      sql.NullTime `json:"test_date"`
	Location      string       `json:"location"`
	Status        string       `json:"status"`
}

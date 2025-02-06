package entities

import (
	"database/sql"
	"encoding/json"
	"time"
)

type TestDriveAppointment struct {
	AppointmentID int          `json:"appointment_id"`
	CarID         int          `json:"car_id"`
	ClientID      int          `json:"client_id"`
	TestDate      sql.NullTime `json:"test_date"`
	Location      string       `json:"location"`
	Status        string       `json:"status"`
}

func (a *TestDriveAppointment) UnmarshalJSON(data []byte) error {
	type Alias TestDriveAppointment
	aux := &struct {
		TestDate string `json:"test_date"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.TestDate != "" {
		parsedTime, err := time.Parse(time.RFC3339, aux.TestDate)
		if err != nil {
			return err
		}
		a.TestDate = sql.NullTime{Time: parsedTime, Valid: true}
	} else {
		a.TestDate = sql.NullTime{Valid: false}
	}

	return nil
}

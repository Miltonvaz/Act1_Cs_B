package application

import (
	"ejercicio1/src/Appointment/domain/entities"
	"ejercicio1/src/Appointment/infraestructure/adapter"
)

type GetNotifications struct {
	notificationAdapter *adapter.NotificationAdapter
}

func NewGetNotifications(notificationAdapter *adapter.NotificationAdapter) *GetNotifications {
	return &GetNotifications{notificationAdapter: notificationAdapter}
}

func (gn *GetNotifications) Execute() <-chan entities.TestDriveAppointment {
	return gn.notificationAdapter.Subscribe()
}

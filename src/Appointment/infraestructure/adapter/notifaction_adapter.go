package adapter

import (
	"ejercicio1/src/Appointment/domain/entities"
	"sync"
)

type NotificationAdapter struct {
	subscribers []chan entities.TestDriveAppointment
	mutex       sync.Mutex
}

func NewNotificationAdapter() *NotificationAdapter {
	return &NotificationAdapter{
		subscribers: make([]chan entities.TestDriveAppointment, 0),
	}
}

func (n *NotificationAdapter) Subscribe() chan entities.TestDriveAppointment {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	ch := make(chan entities.TestDriveAppointment, 1)
	n.subscribers = append(n.subscribers, ch)
	return ch
}

func (n *NotificationAdapter) Notify(appointment entities.TestDriveAppointment) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	for _, subscriber := range n.subscribers {
		select {
		case subscriber <- appointment:
		default:
		}
	}
}

package controllers

import (
	"ejercicio1/src/Appointment/infraestructure/adapter"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type NotificationController struct {
	notificationAdapter *adapter.NotificationAdapter
}

func NewNotificationController(notificationAdapter *adapter.NotificationAdapter) *NotificationController {
	return &NotificationController{notificationAdapter: notificationAdapter}
}

func (nc *NotificationController) Execute(c *gin.Context) {
	subscriber := nc.notificationAdapter.Subscribe()
	timeout := time.After(30 * time.Second)

	select {
	case appointment := <-subscriber:
		fmt.Println("Notificación recibida: ", appointment)
		c.JSON(http.StatusOK, gin.H{
			"appointment": appointment,
		})
	case <-timeout:
	}
}

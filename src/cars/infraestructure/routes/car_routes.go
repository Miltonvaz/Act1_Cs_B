package routes

import (
	"ejercicio1/src/cars/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterCarRoutes(
	r *gin.Engine,
	createCarController *controllers.CreateCarController,
	viewCarController *controllers.ListCarController,
	viewByIdCarController *controllers.ViewByIdCarController,
	updateCarController *controllers.UpdateCarController,
	deleteCarController *controllers.DeleteCarController,
	updateAvailabilityCarController *controllers.UpdateAvailabilityCarController,
	getAvailableCarsController *controllers.GetAvailableCarsController,
) {
	r.POST("/cars", createCarController.Execute)
	r.GET("/cars", viewCarController.Execute)
	r.GET("/cars/:id", viewByIdCarController.Execute)
	r.PUT("/cars/:id", updateCarController.Execute)
	r.DELETE("/cars/:id", deleteCarController.Execute)
	r.PUT("/cars/:id/availability", updateAvailabilityCarController.Execute)
	r.GET("/cars/available", getAvailableCarsController.Execute)
}

package main

import (
	"github.com/gin-contrib/cors"
	"log"

	"ejercicio1/src/Appointment/infraestructure/dependencies_a"
	"ejercicio1/src/Appointment/infraestructure/routes_a"
	"ejercicio1/src/cars/infraestructure/dependencies"
	"ejercicio1/src/cars/infraestructure/routes"
	dependenciesc "ejercicio1/src/clients/infraestructure/dependencies_c"
	routesc "ejercicio1/src/clients/infraestructure/routes_c"

	"github.com/gin-gonic/gin"
)

func main() {

	createClientController, viewClientController, editClientController, deleteClientController, viewByIdClientController, authController, err := dependenciesc.Init()
	if err != nil {
		log.Fatalf("Error initializing client dependencies: %v", err)
		return
	}

	createCarController, viewCarController, deleteCarController, updateCarController, viewByIdCarController, updateAvailabilityCarController, getAvailableCarsController, err := dependencies.InitCars()
	if err != nil {
		log.Fatalf("Error initializing car dependencies: %v", err)
		return
	}

	createAppointmentController, listAppointmentController, deleteAppointmentController, updateAppointmentController, viewByIdAppointmentController, updateAppointmentStatusController, viewAppointmentStatusController, notificationController, err := dependencies_a.InitAppointments()
	if err != nil {
		log.Fatalf("Error initializing appointment dependencies: %v", err)
		return
	}

	r := gin.Default()
	r.Use(cors.Default())

	routes.RegisterCarRoutes(r,
		createCarController,
		viewCarController,
		viewByIdCarController,
		updateCarController,
		deleteCarController,
		updateAvailabilityCarController,
		getAvailableCarsController,
	)

	routesc.RegisterClientRoutes(r,
		createClientController,
		viewClientController,
		editClientController,
		deleteClientController,
		viewByIdClientController,
		authController,
	)

	routes_a.RegisterAppointmentRoutes(r,
		createAppointmentController,
		listAppointmentController,
		viewByIdAppointmentController,
		updateAppointmentController,
		deleteAppointmentController,
		updateAppointmentStatusController,
		viewAppointmentStatusController,
		notificationController,
	)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

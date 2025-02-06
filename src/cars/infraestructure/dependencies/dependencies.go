package dependencies

import (
	"ejercicio1/src/cars/application"
	"ejercicio1/src/cars/infraestructure/controllers"
	"ejercicio1/src/cars/infraestructure/db"
	"ejercicio1/src/core"
)

func InitCars() (
	*controllers.CreateCarController,
	*controllers.ListCarController,
	*controllers.DeleteCarController,
	*controllers.UpdateCarController,
	*controllers.ViewByIdCarController,
	*controllers.UpdateAvailabilityCarController,
	*controllers.GetAvailableCarsController,
	error,
) {
	pool := core.GetDBPool()
	ps := db.NewMySQL(pool.DB)

	createCar := application.NewCreateCar(ps)
	listCar := application.NewListCar(ps)
	updateCar := application.NewUpdateCar(ps)
	deleteCar := application.NewDeleteCar(ps)
	viewByIdCar := application.NewViewByIdCar(ps)
	updateCarAvailability := application.NewUpdateAvailabilityCar(ps)
	getAvailableCars := application.NewGetAvailableCars(ps)

	createCarController := controllers.NewCreateCarController(*createCar)
	viewCarController := controllers.NewListCarController(*listCar)
	updateCarController := controllers.NewUpdateCarController(*updateCar)
	deleteCarController := controllers.NewDeleteCarController(*deleteCar)
	viewByIdCarController := controllers.NewViewByIdCarController(*viewByIdCar)
	updateCarAvailabilityController := controllers.NewUpdateAvailabilityCarController(*updateCarAvailability)
	getAvailableCarsController := controllers.NewGetAvailableCarsController(*getAvailableCars)
	return createCarController, viewCarController, deleteCarController, updateCarController, viewByIdCarController, updateCarAvailabilityController, getAvailableCarsController, nil
}

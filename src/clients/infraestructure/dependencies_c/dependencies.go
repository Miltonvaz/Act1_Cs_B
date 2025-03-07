package dependenciesc

import (
	"ejercicio1/src/clients/application"
	"ejercicio1/src/clients/infraestructure/db"
	"ejercicio1/src/clients/infraestructure/controllers"
	"ejercicio1/src/core"
)

func Init() (
	*controllers.CreateClientController, 
	*controllers.ViewClientController, 
	*controllers.EditClientController, 
	*controllers.DeleteClientController, 
	*controllers.ViewClientByIdController,
	*controllers.AuthController,
	error,
) {
	
	pool := core.GetDBPool()
	ps := db.NewMySQL(pool.DB)

	createClient := application.NewCreateClient(ps)
	viewClient := application.NewListClient(ps) 
	editClient := application.NewEditClient(ps)
	deleteClient := application.NewDeleteClient(ps)
	viewClientById := application.NewClientById(ps)
	authService := application.NewAuthService(ps)

	authController := controllers.NewAuthController(authService)
	createClientController := controllers.NewCreateClientController(*createClient)
	viewClientController := controllers.NewViewClientController(*viewClient)
	editClientController := controllers.NewEditClientController(*editClient)
	deleteClientController := controllers.NewDeleteClientController(*deleteClient)
	viewClientByIdController := controllers.NewViewClientByIdController(*viewClientById)


	return createClientController, viewClientController, editClientController, deleteClientController, viewClientByIdController, authController, nil
}

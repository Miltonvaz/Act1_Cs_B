package routesc

import (
	"ejercicio1/src/clients/infraestructure/controllers"
	"ejercicio1/src/core/security"

	"github.com/gin-gonic/gin"
)

func RegisterClientRoutes(r *gin.Engine, 
	createClientController *controllers.CreateClientController,
	viewClientController *controllers.ViewClientController,
	editClientController *controllers.EditClientController,
	deleteClientController *controllers.DeleteClientController,
	viewByIdClientController *controllers.ViewClientByIdController,
	loginController *controllers.AuthController,
	){
	r.POST("/clients",createClientController.Execute)

	r.GET("/clients",security.JWTMiddleware(), viewClientController.Execute)

	r.POST("/login", loginController.Execute)

	r.GET("/clients/:id", security.JWTMiddleware(),viewByIdClientController.Execute)


	r.PUT("/clients/:id",security.JWTMiddleware(), editClientController.Execute)

	r.DELETE("/clients/:id",security.JWTMiddleware(), deleteClientController.Execute)
}

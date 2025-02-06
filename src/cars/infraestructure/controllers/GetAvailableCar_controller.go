package controllers

import (
	"ejercicio1/src/cars/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAvailableCarsController struct {
	usecase application.GetAvailableCars
}

func NewGetAvailableCarsController(usecase application.GetAvailableCars) *GetAvailableCarsController {
	return &GetAvailableCarsController{usecase: usecase}
}

func (gc *GetAvailableCarsController) Execute(c *gin.Context) {
	cars, err := gc.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve available cars"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"cars": cars,
	})
}

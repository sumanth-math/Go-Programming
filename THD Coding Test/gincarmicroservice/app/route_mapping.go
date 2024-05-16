package app

import (
	"github.com/homedepot/gincarmicroservice/controllers"
)

const (
	carsRoute       string = "/cars"
	getCarByIdRoute string = "/cars/:id"
	addCarRoute     string = "/cars"
)

var (
	carsController = controllers.CarsController
)

func mapRoutes() {
	router.GET(carsRoute, carsController.GetCarsInventory)
	router.GET(getCarByIdRoute, carsController.GetCarDetails)
	router.POST(carsRoute, carsController.AddCarToInventory)
}

package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/homedepot/gincarmicroservice/model/cars"
	"github.com/homedepot/gincarmicroservice/services"
)

type carsController struct{}

const (
	id string = "id"
)

var (
	CarsController = carsController{}
)

func (controller carsController) GetCarsInventory(c *gin.Context) {
	cars, err := services.CarsService.GetAllCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (controller carsController) GetCarDetails(c *gin.Context) {
	carId := c.Param(id)
	car, err := services.CarsService.GetCarDetailsById(carId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if car.ID == carId {
		c.JSON(http.StatusOK, car)
		return
	}

	c.JSON(http.StatusBadRequest,
		fmt.Sprintf("No car with id %s found!", carId))
}

func (controller carsController) AddCarToInventory(c *gin.Context) {
	var car cars.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
		return
	}

	isSuccess, err :=
		services.CarsService.AddCarToInventory(car)
	if err != nil || !isSuccess {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, car)
}

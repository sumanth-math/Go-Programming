package services

import (
	"os"

	"github.com/homedepot/gincarmicroservice/datasources"
	"github.com/homedepot/gincarmicroservice/model/cars"
)

type carsService struct{}

var (
	fileName       = os.Getenv("FILE_NAME")
	fileDataSource = datasources.FileCarDataSource{FileName: fileName}
	//dbDataSource   = datasources.DBCarDataSource{DataSourceName: ""}
	CarsService = carsService{}
)

// GetAllCars retrieves all the cars in the inventory from the
// given (file/db) data source
func (service carsService) GetAllCars() ([]cars.Car, error) {
	// return dbDataSource.GetAllCars()
	return fileDataSource.GetAllCars()
}

// GetCarDetailsById retrieves the car details for the given id from the
// given (file/db) data source
func (service carsService) GetCarDetailsById(id string) (cars.Car, error) {
	// return dbDataSource.GetCarDetailsById(id)
	return fileDataSource.GetCarDetailsById(id)
}

// AddCarToInventory adds the given car details into the
// given (file/db) data source
func (service carsService) AddCarToInventory(car cars.Car) (bool, error) {
	// return dbDataSource.AddNewCarDetails(car)
	return fileDataSource.AddNewCarDetails(car)
}

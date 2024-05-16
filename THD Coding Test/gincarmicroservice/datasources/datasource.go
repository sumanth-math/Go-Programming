package datasources

import (
	"github.com/homedepot/gincarmicroservice/model/cars"
)

type CarDataSource interface {
	GetAllCars() ([]cars.Car, error)
	GetCarDetailsById(id string) (cars.Car, error)
	AddNewCarDetails(car cars.Car) (bool, error)
}

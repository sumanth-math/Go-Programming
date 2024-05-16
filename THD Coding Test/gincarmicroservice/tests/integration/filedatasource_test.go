// Package Integration

package integration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/homedepot/gincarmicroservice/model/cars"
	"github.com/homedepot/gincarmicroservice/services"
)

/************************************
Example Test Command
go test ./... -tags=integration -v
************************************/

func TestGetAllCars(t *testing.T) {
	cars, err := services.CarsService.GetAllCars()
	fmt.Printf("Cars Inventory: %+v\n", cars)
	assert.Nil(
		t,
		err,
		fmt.Sprintf("Get All Cars API error: %s", err),
	)
	assert.Equal(
		t,
		len(cars),
		7,
	)
}

func TestGetCarDetailsById(t *testing.T) {
	car, err := services.CarsService.GetCarDetailsById("PXij34opq")
	fmt.Printf("Cars Inventory: %+v\n", car)
	assert.Nil(
		t,
		err,
		fmt.Sprintf("Get All Cars API error: %s", err),
	)
	assert.Equal(
		t,
		car.ID,
		"PXij34opq",
	)

}

func TestGetCarDetailsByInvalidId(t *testing.T) {
	car, err := services.CarsService.GetCarDetailsById("UUUOO77543")
	fmt.Printf("Cars Inventory: %+v\n", car)
	expectedErrorMessage := fmt.Sprintf("no car found for the requested id %s", car.ID)
	assert.Error(
		t,
		err,
		expectedErrorMessage,
	)
	assert.Equal(
		t,
		car.ID,
		"",
	)

}

func TestAddCarSuccess(t *testing.T) {
	car := cars.Car{
		ID:       "UHFR563klm",
		Make:     "Buick",
		Model:    "Mercury",
		Package:  "SE",
		Color:    "Green",
		Year:     2015,
		Category: "Sport",
		Mileage:  181920,
		Price:    46290.99,
	}
	isSuccess, err := services.CarsService.AddCarToInventory(car)
	assert.Nil(
		t,
		err,
		fmt.Sprintf("Add car to inventory error: %s", err),
	)
	assert.Equal(
		t,
		isSuccess,
		true,
	)
}

func TestAddCarIdAlreadyExists(t *testing.T) {
	car := cars.Car{
		ID:       "WKq563klm",
		Make:     "Buick",
		Model:    "Mercury",
		Package:  "SE",
		Color:    "Green",
		Year:     2015,
		Category: "Sport",
		Mileage:  181920,
		Price:    46290.99,
	}
	expectedErrorMessage := fmt.Sprintf("car with id %s already exists", car.ID)
	isSuccess, err := services.CarsService.AddCarToInventory(car)
	assert.EqualError(
		t,
		err,
		expectedErrorMessage,
	)
	assert.Equal(
		t,
		isSuccess,
		false,
	)
}

func TestAddCarInvalidDetails(t *testing.T) {
	car := cars.Car{
		ID:       "9iosefre",
		Make:     "Buick",
		Model:    "Mercury",
		Package:  "SE",
		Color:    "Green",
		Year:     2015,
		Category: "Sport",
		Mileage:  181920,
		Price:    -23.75,
	}
	isSuccess, err := services.CarsService.AddCarToInventory(car)
	assert.EqualError(
		t,
		err,
		"car price is not valid",
	)
	assert.Equal(
		t,
		isSuccess,
		false,
	)
}

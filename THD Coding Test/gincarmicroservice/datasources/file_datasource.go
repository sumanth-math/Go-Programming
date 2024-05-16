package datasources

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/homedepot/gincarmicroservice/model/cars"
)

type FileCarDataSource struct {
	FileName string
}

// GetAllCars retrieves all the cars in the inventory from the file
func (r FileCarDataSource) GetAllCars() ([]cars.Car, error) {
	var cars []cars.Car

	// Get cars from file
	carsByteArray, err := getfileContentsAsByteArray(r.FileName)

	// Unable to get content from file
	if err != nil || len(carsByteArray) == 0 {
		return cars, err
	}

	// Unmarshal the json file contents to car array
	json.Unmarshal(carsByteArray, &cars)

	return cars, nil
}

// GetCarDetailsById retrieves the car details for the given id from file
func (r FileCarDataSource) GetCarDetailsById(id string) (cars.Car, error) {
	// Get all cars
	inventoryCars, err := r.GetAllCars()
	if err != nil {
		return cars.Car{}, err
	}
	// Loop through cars to find the requested car details
	for i := 0; i < len(inventoryCars); i++ {
		// return the car details for the one that matches the given id
		if strings.EqualFold(inventoryCars[i].ID, id) {
			return inventoryCars[i], nil
		}
	}

	// No car found with the requested id
	return cars.Car{},
		fmt.Errorf("no car found for the requested id %s", id)
}

// AddNewCarDetails adds the given car details into file
func (r FileCarDataSource) AddNewCarDetails(car cars.Car) (bool, error) {
	// check car data to ensure everything is valid
	if err := car.Validate(); err != nil {
		return false, err
	}
	// Check if id of the car to add already exists
	cars, err := r.GetAllCars()
	if err != nil {
		return false, err
	}

	// Check if the given id alreadys exists
	for i := 0; i < len(cars); i++ {
		if strings.EqualFold(cars[i].ID, car.ID) {
			return false,
				fmt.Errorf("car with id %s already exists", car.ID)
		}
	}

	// Append new car detail to the car inventory list
	cars = append(cars, car)
	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(cars)
	if err != nil {
		return false,
			errors.New("error while marshaling data")
	}

	// Update the file with new car details
	err = ioutil.WriteFile(r.FileName, dataBytes, 0644)
	if err != nil {
		return false,
			errors.New("error while writing to the file data source")
	}

	return true, nil
}

func getfileContentsAsByteArray(fileName string) ([]byte, error) {
	// open cars json file
	jsonCarFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File ERROR: ", err.Error())
		return []byte{}, err
	}

	// close the file when done
	defer jsonCarFile.Close()

	// retrieve file contents
	carsByteValue, _ := ioutil.ReadAll(jsonCarFile)

	return carsByteValue, nil
}

package datasources

import (
	"database/sql"

	"github.com/homedepot/gincarmicroservice/model/cars"
)

const (
	driverName            = "mysql"
	getAllCarsProc        = "CALL Car_GetAll()"
	getCarDetailsByIdProc = "CALL Car_GetById(?)"
	addCarDetailsProc     = "CALL CustomerInfo_Upsert_PreferredBeAnyWhere(?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

type DBCarDataSource struct {
	DataSourceName string
}

// AddNewCarDetails retrives cars inventory data from DB
func (r DBCarDataSource) GetAllCars() ([]cars.Car, error) {
	db, err := sql.Open(driverName, r.DataSourceName)
	if err != nil {
		return []cars.Car{}, err
	}

	// close db
	defer db.Close()
	rows, err := db.Query(getAllCarsProc)
	if err != nil {
		return []cars.Car{}, err
	}

	defer rows.Close()
	var inventoryCars []cars.Car
	for rows.Next() {
		var car cars.Car
		scanErr := rows.Scan(&car.ID, &car.Make, &car.Model,
			&car.Package, &car.Color, &car.Year,
			&car.Category, &car.Mileage, &car.Price)
		if scanErr != nil { // error during scan
			return []cars.Car{}, scanErr
		}
		inventoryCars = append(inventoryCars, car)
	}

	return inventoryCars, nil
}

// Retrives car details for the given id from DB
func (r DBCarDataSource) GetCarDetailsById(id string) (cars.Car, error) {
	db, err := sql.Open("driverName", r.DataSourceName)
	if err != nil {
		return cars.Car{}, err
	}

	// close db
	defer db.Close()
	stmt, err := db.Query(getCarDetailsByIdProc, id)
	if err != nil {
		return cars.Car{}, err
	}
	defer stmt.Close()
	stmt.Next()
	var car cars.Car
	scanErr := stmt.Scan(&car.ID, &car.Make, &car.Model,
		&car.Package, &car.Color, &car.Year,
		&car.Category, &car.Mileage, &car.Price)

	if scanErr != nil {
		return cars.Car{}, err
	}

	return car, nil
}

// Add the given car details into DB
func (r DBCarDataSource) AddNewCarDetails(car cars.Car) (bool, error) {
	db, err := sql.Open("driverName", r.DataSourceName)
	if err != nil {
		return false, err
	}

	// close db
	defer db.Close()
	result, err := db.Exec(addCarDetailsProc, car.ID, car.Make,
		car.Model, car.Package, car.Color, car.Year, car.Category,
		car.Mileage, car.Price)
	if err != nil {
		return false, err
	}

	_, rowErr := result.RowsAffected()
	if rowErr != nil {
		return false, rowErr
	}

	return true, nil
}

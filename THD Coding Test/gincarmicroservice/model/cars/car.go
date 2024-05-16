package cars

import (
	"errors"
	"time"
)

const(
	emptyString string = ""
	maxCarAge int = 20
	maxCarMileage int = 200000
	minCarPrice float64 = 500.00
)

// Car Data Model
type Car struct {
	ID       string  `json:"id"`
	Make     string  `json:"make"`
	Model    string  `json:"model"`
	Package  string  `json:"package"`
	Color    string  `json:"color"`
	Year     int  `json:"year"`
	Category string  `json:"category"`
	Mileage  int     `json:"mileage"`
	Price    float64 `json:"price"`
}

func (c Car) Validate() error {
	// Invlaid ID
	if c.ID == emptyString {
		return errors.New("invalid car id")
	}
	// Invalid Make
	if c.Make == emptyString {
		return errors.New("invalid car make")
	}
	// Invalid Model
	if c.Model == emptyString {
		return errors.New("invalid car model")
	}
	// Invalid Color
	if c.Color == emptyString {
		return errors.New("invalid car color")
	}
	// Invalid Year
	currentYear := time.Now().Year()
	if c.Year <= 0 || c.Year > currentYear ||
	 c.Year <= currentYear - maxCarAge {
		return errors.New("car too old for inventory")
	}
	// Invalid Category
	if c.Category == emptyString {
		return errors.New("invalid car category")
	}
	// Invalid Car mileage
	if c.Mileage < 0 || c.Mileage > maxCarMileage {
		return errors.New("car has lot of mileage for inventory")
	}
	// Invalid car price
	if c.Price < minCarPrice{
		return errors.New("car price is not valid")
	}

	// All car details valid
	return nil
}

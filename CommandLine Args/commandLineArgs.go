package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// First argument is the path to the executable which can be excluded
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Command Line Argument %d: %s\n", i, os.Args[i])
		number, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println("Invalid argument for factorial calculation. Please provide a valid integer.")
		} else {
			factorialValue, err := Factorial(number)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Factorial value of %d: %d\n", number, factorialValue)
			}
		}
	}
}

func Factorial(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("Factorial of negative number is not defined")
	} else {
		return factorialCalculator(n), nil
	}
}

func factorialCalculator(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorialCalculator(n-1)
}

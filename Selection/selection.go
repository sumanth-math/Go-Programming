package main

import (
	"errors"
	"fmt"
)

func main() {
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15}
	sumMatch := 25

	triplets, err := getTripletsMatchingSum(inputList, sumMatch)

	if err != nil {
		fmt.Println(err)
	} else {
		for _, triplet := range triplets {
			fmt.Println(triplet)
		}
	}
}

func getTripletsMatchingSum(inputList []int, sumMatch int) ([][3]int, error) {
	var triplets [][3]int
	if len(inputList) < 3 {
		return triplets, errors.New("List should have atleast 3 numbers")
	}
	duplicateEntry := make(map[int]bool)

	for i := 0; i < len(inputList); i++ {
		j := i + 1
		if j == len(inputList) {
			j = 0
		}
		for {
			k := j + 1
			if k == len(inputList) {
				k = 0
			}

			if inputList[i]+inputList[j]+inputList[k] == sumMatch {
				if _, ok := duplicateEntry[(i+1)*(j+1)*(k+1)]; !ok {
					duplicateEntry[(i+1)*(j+1)*(k+1)] = true
					var triplet [3]int
					triplet[0] = inputList[i]
					triplet[1] = inputList[j]
					triplet[2] = inputList[k]
					triplets = append(triplets, triplet)
				}
			}
			j++

			if j == len(inputList) {
				j = 0
			}
			if j == i || j+1 == i {
				break
			}
		}
	}

	return triplets, nil
}

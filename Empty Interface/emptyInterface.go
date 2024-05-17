package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(evaluateInput(1, 2))
	fmt.Println(evaluateInput("Hello", "World"))

}

func evaluateInput(input1 interface{}, input2 interface{}) interface{} {
	if reflect.TypeOf(input1).String() == "int" &&
		reflect.TypeOf(input2).String() == "int" {
		return input1.(int) + input2.(int)
	} else if reflect.TypeOf(input1).String() == "string" &&
		reflect.TypeOf(input2).String() == "string" {
		return input1.(string) + " " + input2.(string)
	}

	return nil
}

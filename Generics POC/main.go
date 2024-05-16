package main

import (
	"fmt"
	"genericpoc/integer"
)

var intLinkedList *integer.Node

func main() {
	fmt.Println("Linked List")
	integer.AddNode(&intLinkedList, 10)
	integer.AddNode(&intLinkedList, 100)
	integer.AddNode(&intLinkedList, 1000)
	integer.DisplayLinkedList(intLinkedList)
	integer.DeleteNode(&intLinkedList, 123)
	integer.DeleteNode(&intLinkedList, 10)
	integer.DeleteNode(&intLinkedList, 100)
	integer.DeleteNode(&intLinkedList, 1000)

	integer.DisplayLinkedList(intLinkedList)
}

package integer

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func AddNode(linkedList **Node, data int) {
	if *linkedList == nil {
		newNode := new(Node)
		*&newNode.Data = data
		*linkedList = newNode
		return
	}

	node := *linkedList
	for {
		if node.Next == nil {
			break
		}
		node = node.Next
	}

	newNode := new(Node)
	*&newNode.Data = data
	node.Next = newNode
}

func DeleteNode(linkedList **Node, data int) {
	if *linkedList == nil {
		fmt.Printf("Linked list empty. Nothing to delete!")
		return
	}

	currentNode := *linkedList
	previousNode := *linkedList
	nodeFound := false

	for {
		if currentNode.Data == data {
			nodeFound = true
			break
		} else if currentNode.Next == nil {
			break
		} else {
			previousNode = currentNode
			currentNode = currentNode.Next
		}
	}
	if nodeFound {
		// First Node in the list
		if currentNode == previousNode {
			if currentNode.Next == nil {
				*linkedList = nil
			} else {
				*linkedList = currentNode.Next
			}
		} else {
			previousNode.Next = currentNode.Next
		}
	} else {
		fmt.Printf("Not node found with data value: %d\n", data)
	}
}

func DisplayLinkedList(linkedList *Node) {
	if linkedList == nil {
		fmt.Printf("No element in the list!")
		return
	}
	node := linkedList
	for {
		if node.Next == nil {
			fmt.Printf("Value=%d\n", node.Data)
			break
		} else {
			fmt.Printf("Value=%d\n", node.Data)
			node = node.Next
		}
	}
}

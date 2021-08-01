package main

import (
	"fmt"

	linkedList "github.com/dan-almenar/golang/datastructures/linkedlist"
)

func main() {
	randomString := "some random string"
	single := linkedList.NewSingleLinkedList()
	single.AddNode(randomString)
	single.AddNode(5)
	single.AddNode(false)
	single.AddNode([]byte("random byte string"))
	single.AddNode(3.14159)
	fmt.Println(single.ListData())

	//since nodes are pointers, their data can be easily replaced if they exist,
	//thus, avoiding the need to delete -> insert
	isNode, node := single.FindNode(5)
	if isNode {
		node.Data = 13
	}
	fmt.Println("isNode value:", isNode)

	// deleting and inserting nodes:
	for any, ok := range single.ListData() {
		fmt.Println(any, ok)
	}
	single.DeleteNode("some random string")

	//looping over the Linked List:
	for any, ok := range single.ListData() {
		fmt.Println(any, ok)
	}
	single.InsertNode(false, "inserted node")
	for any, ok := range single.ListData() {
		fmt.Println(any, ok)
	}

	doubly := linkedList.NewDoublyLinkedList()
	doubly.AddNode(true)
	doubly.AddNode(15)
	doubly.AddNode(30)
	fmt.Println(doubly.ListData())
	doubly.InsertNode(30, false)
	fmt.Println(doubly.ListData())
	isDoublyNode, nodeVal := doubly.FindNode(false)
	fmt.Println(isDoublyNode, nodeVal.PreviousNode, nodeVal.NextNode)
	doubly.DeleteNode(false)
	fmt.Println(doubly.ListData())

	//replacing node.Data value
	isDoublyNode, nodeVal = doubly.FindNode(15)
	if isDoublyNode {
		nodeVal.Data = []int{10, 20, 30}
	}
	fmt.Println(doubly.ListData())

	//head node.PreviousNode = nil
	isDoublyNode, nodeVal = doubly.FindNode(30)
	if isDoublyNode {
		fmt.Println(nodeVal.PreviousNode)
	}
}

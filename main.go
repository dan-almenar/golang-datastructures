package main

import (
	"fmt"

	linkedList "github.com/dan-almenar/golang/datastructures/linkedlist"
)

func main() {
	randomString := "some random string"
	list := linkedList.NewSingleLinkedList()
	list.AddNode(randomString)
	list.AddNode(5)
	list.AddNode(false)
	list.AddNode([]byte("random byte string"))
	list.AddNode(3.14159)
	fmt.Println(list.ListData())

	//since nodes are pointers, their data can be easily replaced if they exist,
	//thus, avoiding the need to delete -> insert
	isNode, node := list.FindNode(5)
	if isNode {
		node.Data = 13
	}
	fmt.Println("isNode value:", isNode)

	// deleting and inserting nodes:
	for any, ok := range list.ListData() {
		fmt.Println(any, ok)
	}
	list.DeleteNode("some random string")
	for any, ok := range list.ListData() {
		fmt.Println(any, ok)
	}
	list.InsertNode(false, "inserted node")
	for any, ok := range list.ListData() {
		fmt.Println(any, ok)
	}
}

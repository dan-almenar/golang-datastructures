package linkedList

import (
	"fmt"
)

type any interface{}

type node struct {
	Data     any
	NextNode *node
}

type SingleLinkedList struct {
	Head   *node
	Lenght int
}

func newNode(d any) *node {
	return &node{
		Data: d,
	}
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		Lenght: 0,
	}
}

func (l *SingleLinkedList) AddNode(d any) {
	node := newNode(d)
	if l.Lenght == 0 {
		l.Head = node
	}
	newNext := l.Head
	l.Head = node
	l.Head.NextNode = newNext
	l.Lenght++
}

func (l SingleLinkedList) ListData() []any {
	nodeToPrint := l.Head
	data := []any{}
	for l.Lenght != 0 {
		data = append(data, nodeToPrint.Data)
		nodeToPrint = nodeToPrint.NextNode
		l.Lenght--
	}
	return data
}

func (l SingleLinkedList) findPreviousNode(v any) *node {
	previousNode := l.Head
	for l.Lenght != 0 {
		if previousNode.NextNode.Data == v {
			return previousNode
		}
		previousNode = previousNode.NextNode
		l.Lenght--
	}
	return nil
}

func (l SingleLinkedList) FindNode(v any) (bool, *node) {
	nodeToVerify := l.Head
	for l.Lenght != 0 {
		if nodeToVerify.Data == v {
			return true, nodeToVerify
		}
		nodeToVerify = nodeToVerify.NextNode
		l.Lenght--
	}
	return false, nil
}

func (l *SingleLinkedList) DeleteNode(v any) error {
	previousNode := l.findPreviousNode(v)
	if previousNode == nil {
		err := fmt.Errorf("value is noton the list")
		return err
	}
	previousNode.NextNode = previousNode.NextNode.NextNode
	fmt.Println("Successfully deleted node with value:", v)
	l.Lenght--
	return nil
}

// inserts a new node with Data = v after a node with Data = nodeData.
// if no node holds nodeData as Data, it appends the new node with the addNode method.
func (l *SingleLinkedList) InsertNode(nodeData, v any) {
	_, previousNode := l.FindNode(nodeData)
	if previousNode == nil {
		fmt.Println("Value was not on the list. New node appended as Head with Data:", nodeData)
		l.AddNode(v)
		return
	}
	node := newNode(v)
	node.NextNode = previousNode.NextNode
	previousNode.NextNode = node
	l.Lenght++
}

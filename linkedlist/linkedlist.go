package linkedList

import (
	"fmt"
)

type any interface{}

type singleNode struct {
	Data     any
	NextNode *singleNode
}

type doublyNode struct {
	Data         any
	NextNode     *doublyNode
	PreviousNode *doublyNode
}

type SingleLinkedList struct {
	Head   *singleNode
	Lenght int
}

type DoublyLinkedList struct {
	Head   *doublyNode
	Lenght int
}

func newSingleNode(d any) *singleNode {
	return &singleNode{
		Data: d,
	}
}

func newDoublyNode(d any) *doublyNode {
	return &doublyNode{
		Data: d,
	}
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		Lenght: 0,
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Lenght: 0,
	}
}

func (l *SingleLinkedList) AddNode(d any) {
	node := newSingleNode(d)
	if l.Lenght == 0 {
		l.Head = node
	}
	newNext := l.Head
	l.Head = node
	l.Head.NextNode = newNext
	l.Lenght++
}

func (l *DoublyLinkedList) AddNode(d any) {
	node := newDoublyNode(d)
	if l.Lenght == 0 {
		l.Head = node
	}
	l.Head.PreviousNode = node
	newNext := l.Head
	l.Head = node
	l.Head.NextNode = newNext
	newNext.PreviousNode = l.Head
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

func (l DoublyLinkedList) ListData() []any {
	nodeToPrint := l.Head
	data := []any{}
	for l.Lenght != 0 {
		data = append(data, nodeToPrint.Data)
		nodeToPrint = nodeToPrint.NextNode
		l.Lenght--
	}
	return data
}

func (l SingleLinkedList) findPreviousNode(v any) *singleNode {
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

func (l SingleLinkedList) FindNode(v any) (bool, *singleNode) {
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

func (l DoublyLinkedList) FindNode(v any) (bool, *doublyNode) {
	nodeToVerify := l.Head
	for l.Lenght != 0 {
		if nodeToVerify.Data == v {
			return true, nodeToVerify
		}
		nodeToVerify = nodeToVerify.NextNode
		l.Lenght--
	}
	nodeToVerify = nodeToVerify.NextNode
	l.Lenght--
	return false, nil
}

func (l *SingleLinkedList) DeleteNode(v any) error {
	previousNode := l.findPreviousNode(v)
	if previousNode == nil {
		err := fmt.Errorf("value is not on the list")
		return err
	}
	previousNode.NextNode = previousNode.NextNode.NextNode
	fmt.Println("Successfully deleted node with value:", v)
	l.Lenght--
	return nil
}

func (l *DoublyLinkedList) DeleteNode(v any) error {
	_, node := l.FindNode(v)
	if node == nil {
		err := fmt.Errorf("no node holds the value %v", v)
		return err
	}
	previous := node.PreviousNode
	next := node.NextNode
	previous.NextNode = next
	next.PreviousNode = previous
	l.Lenght--
	fmt.Println("Successfully deleted node with value:", v)
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
	node := newSingleNode(v)
	node.NextNode = previousNode.NextNode
	previousNode.NextNode = node
	l.Lenght++
}

func (l *DoublyLinkedList) InsertNode(nodeData, v any) {
	_, node := l.FindNode(nodeData)
	if node == nil {
		fmt.Printf("No node holds the value %v. Appended a new node as Head.\n", nodeData)
		l.AddNode(v)
		l.Lenght++
		return
	}
	newNode := newDoublyNode(v)
	previousNext := node.NextNode
	previousNext.PreviousNode = newNode
	node.NextNode = newNode
	newNode.NextNode = previousNext
	newNode.PreviousNode = node
	l.Lenght++
}

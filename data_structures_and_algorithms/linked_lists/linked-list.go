package main

import (
	"fmt"
)

type Node struct {
	next *Node
	data interface{}
}

type LinkedList struct {
	head *Node
}

// no error checking for index
func (list *LinkedList) Read(index int) interface{} {
	currentNode := list.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data
}

func (list *LinkedList) IndexOf(arg interface{}) int {
	currentNode := list.head
	index := 0
	for currentNode != nil {
		if currentNode.data == arg {
			return index
		}
		index += 1
		currentNode = currentNode.next
	}

	return -1

}

func (list *LinkedList) InsertAtIndex(index int, arg interface{}) {
	var prevNode *Node
	currentNode := list.head

	for i := 0; i < index; i++ {
		prevNode = currentNode
		currentNode = currentNode.next
	}

	newNode := &Node{currentNode, arg}

	if prevNode != nil {
		prevNode.next = newNode
	}

	if currentNode == list.head {
		list.head = newNode
	}

}

func (list *LinkedList) DeleteAtIndex(index int) {
	var prevNode *Node
	currentNode := list.head

	if index == 0 {
		list.head = currentNode.next
		return
	}

	for i := 0; i < index; i++ {
		prevNode = currentNode
		currentNode = currentNode.next
	}

	prevNode.next = currentNode.next
}

func (list *LinkedList) PrintAll() {
	currentNode := list.head

	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (list *LinkedList) LastElement() interface{} {
	currentNode := list.head

	if list.IsEmpty() {
		return nil
	}

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	return currentNode.data
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) Reverse() {
	var prevNode *Node
	currentNode := list.head
	var nextNode *Node

	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	list.head = prevNode
}

func NewLinkedList(data interface{}) *LinkedList {
	newHead := Node{nil, data}
	newList := LinkedList{&newHead}
	return &newList
}

func main() {
	list := NewLinkedList("oh")
/*
	list.InsertAtIndex(1, "world")
	list.InsertAtIndex(1, "new")
	list.InsertAtIndex(1, "brave")
	list.InsertAtIndex(0, "that")
*/
//	list.DeleteAtIndex(0)
	//	fmt.Printf("%T\n", list.LastElement())

//	list.PrintAll()


	list.Reverse()
	list.PrintAll()

	//	list.InsertAtIndex(0, "oh")

	//	fmt.Println(list.Read(0))
	//	fmt.Println(list.Read(1))
	//	fmt.Println(list.Read(2))
	//	fmt.Println(list.Read(3))
	//	fmt.Println(list.Read(4))

	//	fmt.Print(list.IndexOf("oh"))

}


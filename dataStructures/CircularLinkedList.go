package dataStructures

import "fmt"

type CircularLinkedList[K int | string] struct {
	head *Node[K]
	tail *Node[K]
}

func (list *CircularLinkedList[string]) InsertBeginning(n *Node[string]) {
	if list.head == nil {
		list.head = n
		list.tail = list.head
		list.head.next = list.head
	} else {
		n.next = list.head
		list.head = n
		list.tail.next = list.head
	}
}

func (list *CircularLinkedList[string]) InsertEnd(n *Node[string]) {
	if list.head == nil {
		list.InsertBeginning(n)
	} else {
		n.next = list.head
		list.tail.next = n
		list.tail = n
	}
}

func (list *CircularLinkedList[string]) RemoveBeginning() {
	if list.head == nil {
		fmt.Printf("The list is empty")
	} else {
		list.head = list.head.next
		list.tail.next = list.head
	}
}

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("Empty list")
	} else {
		current := list.head

		fmt.Println(current.data)

		for current.next != nil {
			current = current.next

			fmt.Println(current.data)
		}
	}
}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	list.Display()
}

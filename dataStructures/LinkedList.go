package dataStructures

import "fmt"

type Node[K int | string] struct {
	data K
	next *Node[K]
}

type LinkedList[K int | string] struct {
	head *Node[K]
}

// O(1)
func (list *LinkedList[int]) Insert(data int) {
	newNode := &Node[int]{data: data}

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

// O(n)
func (list *LinkedList[int]) Traverse() {
	current := list.head

	fmt.Println(current.data)

	for current.next != nil {
		current = current.next

		fmt.Println(current.data)
	}
}

// O(n)
func (list *LinkedList[int]) Search(item int) {
	current := list.head
	idx := 0

	if current == nil {
		fmt.Println("Not found")
		return
	}

	for current.next != nil {
		if current.data == item {
			fmt.Printf("Item found at index %v\n", idx)
			return
		}
		idx += 1
		current = current.next
	}

	fmt.Println("Not found")
}

// O(n)
func (list *LinkedList[int]) Delete(item int) {
	current := list.head

	if current == nil {
		return
	}

	if current.data == item {
		list.head = current.next
		return
	}

	for current.next != nil {
		if current.next.data == item {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// func main() {
// 	list := LinkedList{}

// 	list.Insert(10)
// 	list.Insert(20)
// 	list.Insert(30)
// 	list.Insert(40)

// 	// list.Traverse()

// 	list.Search(20)

// 	// list.Delete(10)

// 	list.Traverse()
// }

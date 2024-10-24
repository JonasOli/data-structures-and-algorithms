package dataStructures

func enqueue(queue CircularLinkedList[int], element int) CircularLinkedList[int] {
	queue.InsertEnd(&Node[int]{data: element})

	return queue
}

func dequeue(queue CircularLinkedList[int]) {
	queue.RemoveBeginning()
}

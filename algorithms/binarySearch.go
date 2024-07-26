package main

import "fmt"

// time complexity = O(log n)
func binarySearch(arr [10]int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		half := (low + high) / 2
		guess := arr[half]

		if guess == item {
			return half
		}
		if guess > item {
			high = half - 1
		} else {
			low = half + 1
		}
	}

	return -1
}

func main() {
	intArr := [10]int{1, 2, 3, 6, 8, 9, 21, 89, 100, 123}

	var itemToSearch int

	fmt.Printf("Array to search: %v\n", intArr)
	fmt.Printf("Enter a number to search on the array: ")
	fmt.Scanln(&itemToSearch)

	res := binarySearch(intArr, itemToSearch)

	if res == -1 {
		fmt.Printf("The result %v was not found on the array!\n", itemToSearch)
	} else {
		fmt.Printf("The result %v was found at the index %v!\n", itemToSearch, res)
	}
}

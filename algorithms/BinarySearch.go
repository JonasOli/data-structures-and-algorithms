package algorithms

// time complexity = O(log n)
func BinarySearch(arr [10]int, item int) int {
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

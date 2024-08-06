package algorithms

import (
	"math/rand/v2"

	"github.com/JonasOli/data-structures-and-algorithms/utils"
)

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// If you select a random pivot, quicksort will be completed in an average time complexity of O(n log n)
	var randomIdx int = rand.IntN(len(arr))
	var pivot int = arr[randomIdx]
	arr = utils.RemoveItemFromList(arr, randomIdx)
	var smaller []int
	var bigger []int

	for _, value := range arr {
		if value > pivot {
			bigger = append(bigger, value)
		} else {
			smaller = append(smaller, value)
		}
	}

	newArr := append([]int{}, QuickSort(smaller)...)
	newArr = append(newArr, pivot)
	newArr = append(newArr, QuickSort(bigger)...)

	return newArr
}

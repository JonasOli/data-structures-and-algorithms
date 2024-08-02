package algorithms

import "github.com/JonasOli/data-structures-and-algorithms/utils"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var arrayHalfIdx int = len(arr) / 2
	var pivot int = arr[arrayHalfIdx]
	arr = utils.RemoveItemFromList(arr, arrayHalfIdx)
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

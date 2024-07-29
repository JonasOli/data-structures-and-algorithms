package tests

import (
	"github.com/JonasOli/data-structures-and-algorithms/algorithms"
	"testing"
)

func Test_ShouldReturnTheCorrectIndexIfTheNumberExistsInTheList(t *testing.T) {
	intArr := [10]int{1, 2, 3, 6, 8, 9, 21, 89, 100, 123}

	itemToSearch := 6
	expectedResult := 3

	var res int = algorithms.BinarySearch(intArr, itemToSearch)

	if res != expectedResult {
		t.Errorf("The result %v is incorrect. Expected: %v", res, expectedResult)
	}
}

func Test_ShouldReturnNegativeOneIfValueNotOnTheList(t *testing.T) {
	intArr := [10]int{1, 2, 3, 6, 8, 9, 21, 89, 100, 123}

	itemToSearch := 987
	expectedResult := -1

	var res int = algorithms.BinarySearch(intArr, itemToSearch)

	if res != expectedResult {
		t.Errorf("The result %v is incorrect. Expected: %v", res, expectedResult)
	}
}

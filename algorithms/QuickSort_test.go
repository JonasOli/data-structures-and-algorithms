package algorithms

import (
	"reflect"
	"testing"
)

func Test_ShouldReturnTheArrayCorrectlyOrdered(t *testing.T) {
	intArr := []int{9, 3, 5, 1, 3, 6, 2}

	newArr := QuickSort(intArr)
	expected := []int{1, 2, 3, 3, 5, 6, 9}

	if !reflect.DeepEqual(newArr, expected) {
		t.Errorf("The array %v is different from the expected %v", newArr, expected)
	}
}

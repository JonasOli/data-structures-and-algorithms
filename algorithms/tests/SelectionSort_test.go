package tests

import (
	"testing"

	"github.com/JonasOli/data-structures-and-algorithms/algorithms"
)

func Test_ShouldReturnTheCorrectElementOrdered(t *testing.T) {
	playListArr := []algorithms.PlayList{
		{ArtistName: "AC/DC", PlayCount: 111},
		{ArtistName: "Metallica", PlayCount: 153},
		{ArtistName: "Iron Maiden", PlayCount: 141},
	}

	newArr := algorithms.SelectionSort(playListArr)
	playListOrderedExpected := []algorithms.PlayList{
		{ArtistName: "Metallica", PlayCount: 153},
		{ArtistName: "Iron Maiden", PlayCount: 141},
		{ArtistName: "AC/DC", PlayCount: 111},
	}

	for idx := range newArr {
		if newArr[idx] != playListOrderedExpected[idx] {
			t.Errorf("The result %v is incorrect. Expected: %v", newArr[idx], playListOrderedExpected[idx])
		}
	}
}

func Test_ShouldReturnAnEmptyArrayIfSendingEmptyList(t *testing.T) {
	playListArr := []algorithms.PlayList{}

	newArr := algorithms.SelectionSort(playListArr)

	if len(newArr) > 0 {
		t.Errorf("The result %v is incorrect. Expected: %v", len(newArr), 0)
	}
}

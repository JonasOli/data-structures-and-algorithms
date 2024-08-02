package algorithms

import "github.com/JonasOli/data-structures-and-algorithms/utils"

type PlayList struct {
	ArtistName string
	PlayCount  int
}

func SelectionSort(list []PlayList) []PlayList {
	newArray := []PlayList{}

	var listLength int = len(list)

	for i := 0; i < listLength; i++ {
		biggerIndex := getBiggerValueIndex(&list)

		newArray = append(newArray, list[biggerIndex])
		list = utils.RemoveItemFromList(list, biggerIndex)
	}

	return newArray
}

func getBiggerValueIndex(list *[]PlayList) int {
	var firstPlayList PlayList = (*list)[0]
	biggerIndex := 0

	for idx := range *list {
		if (*list)[idx].PlayCount > firstPlayList.PlayCount {
			firstPlayList = (*list)[idx]
			biggerIndex = idx
		}
	}

	return biggerIndex
}

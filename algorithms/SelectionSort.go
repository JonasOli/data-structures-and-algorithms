package algorithms

type PlayList struct {
	ArtistName string
	PlayCount  int
}

func SelectionSort(list []PlayList) []PlayList {
	newArray := []PlayList{}

	var listLength int = len(list)

	for i := 0; i < listLength; i++ {
		biggerIndex := getBiggerValue(&list)

		newArray = append(newArray, list[biggerIndex])
		list = remove(list, biggerIndex)
	}

	return newArray
}

func getBiggerValue(list *[]PlayList) int {
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

func remove(list []PlayList, idx int) []PlayList {
	list[idx] = list[len(list)-1]
	return list[:len(list)-1]
}

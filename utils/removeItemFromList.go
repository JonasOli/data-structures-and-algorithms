package utils

func RemoveItemFromList[T comparable](list []T, idx int) []T {
	list[idx] = list[len(list)-1]

	return list[:len(list)-1]
}

package l117

func BinarySearch(arr []int, elem int) int {
	mid := len(arr) / 2

	if elem == arr[mid] {
		return mid
	}
	if elem < arr[mid] {
		left := arr[:mid]
		return BinarySearch(left, elem)
	} else if elem > arr[mid] {
		right := arr[mid:]
		return BinarySearch(right, elem)
	}
	return -1
}

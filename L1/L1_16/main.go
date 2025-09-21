package main

import "fmt"

func main() {
	arr := []int{4, 5, 3, 1, 8, 2, 9, 6, 7, 1120, 12}
	sortedArr := quickSort(arr)

	fmt.Println("arr:", arr)
	fmt.Println("srtd arr:", sortedArr)
}

func quickSort(data []int) []int {
	arr := make([]int, len(data))
	copy(arr, data)

	if len(arr) <= 1 {
		return arr
	}

	pivotId := 0
	pivot := arr[pivotId]

	left := 1
	right := len(arr) - 1

	for left <= right {
		for left <= right && arr[left] <= pivot {
			left++
		}

		for left <= right && arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	arr[pivotId], arr[right] = arr[right], arr[pivotId]

	quickSort(arr[:right])
	quickSort(arr[right+1:])
	return arr
}

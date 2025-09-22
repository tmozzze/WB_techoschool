package l116

func QuickSort(data []int) []int {
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

	QuickSort(arr[:right])
	QuickSort(arr[right+1:])
	return arr
}

package l123

import "log"

func Remove(arr []int, idx int) ([]int, bool) {
	if idx < 0 || idx >= len(arr) {
		log.Printf("index %d out of range", idx)
		return nil, false
	}

	result := make([]int, len(arr)-1)

	copy(result, arr[:idx])
	copy(result[idx:], arr[idx+1:])
	return result, true
}

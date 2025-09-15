package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 3}
	a2 := []int{2, 3, 4, 4}

	result := CalcRepetitions(a1, a2)
	fmt.Println(result)
}

func CalcRepetitions(arr1, arr2 []int) []int {
	m := make(map[int]bool)
	var result []int

	for _, elem := range arr1 {
		m[elem] = true
	}

	for _, elem := range arr2 {
		if m[elem] {
			result = append(result, elem)
			m[elem] = false
		}
	}
	return result
}

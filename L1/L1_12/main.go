package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(UniqueSet(arr))
}

func UniqueSet(arr []string) []string {
	var result []string
	m := make(map[string]bool)

	for _, elem := range arr {
		if !m[elem] {
			m[elem] = true
			result = append(result, elem)
		}
		continue
	}

	return result
}

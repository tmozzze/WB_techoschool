package main

import (
	"fmt"

	l123 "github.com/tmozzze/WB_techoschool/L1/L1_23"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result, ok := l123.Remove(arr, 0)
	if !ok {
		fmt.Println("Remove error")
		return
	}
	fmt.Println("Start arr:", arr)
	fmt.Println("Result arr:", result)

}

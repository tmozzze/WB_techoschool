package main

import (
	"fmt"

	l116 "github.com/tmozzze/WB_techoschool/L1/L1_16"
	l117 "github.com/tmozzze/WB_techoschool/L1/L1_17"
)

func main() {
	arr := []int{2, 4, 3, 10, 1, 17, 9, 1000}

	srtdArr := l116.QuickSort(arr)
	fmt.Println(srtdArr)
	fmt.Println(l117.BinarySearch(srtdArr, 0))
}

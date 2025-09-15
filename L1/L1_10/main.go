package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	Dia := SplitDiapozons(arr)
	fmt.Println(Dia)
}

func SplitDiapozons(arr []float64) map[int][]float64 {
	Diapozons := make(map[int][]float64)

	for _, value := range arr {
		key := int(value/10) * 10
		Diapozons[key] = append(Diapozons[key], value)
	}

	return Diapozons
}

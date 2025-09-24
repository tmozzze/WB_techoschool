package main

import (
	"fmt"

	l120 "github.com/tmozzze/WB_techoschool/L1/L1_20"
)

func main() {
	examples := []string{
		"snow dog sun",
		"hello world",
		"a b c",
		"single",
		"",
		"   multiple   spaces   ",
		" 3",
	}

	for _, example := range examples {
		result := l120.ShaffleWords(example)
		fmt.Printf("Вход: «%s» -> Выход: «%s»\n", example, result)
	}
}

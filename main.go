package main

import (
	"fmt"

	"github.com/tmozzze/WB_techoschool/L2/L2_11/anagrams"
)

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	result := anagrams.FindAnagrams(words)
	fmt.Println(result)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	wg.Add(5)

	for _, num := range arr {
		go func(num int) {
			result := num * num
			fmt.Println(result)
			wg.Done()
		}(num)
	}
	wg.Wait()

}

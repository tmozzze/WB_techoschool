package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := arrToChan(arr)
	out := squareChan(ch)

	for num := range out {
		fmt.Printf("Square num: %d\n", num)
		time.Sleep(500 * time.Millisecond)
	}

}

func arrToChan(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)

	}()
	return out
}

func squareChan(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

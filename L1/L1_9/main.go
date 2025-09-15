package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer log.Println("main finish")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := arrToChan(arr)
	result := squareChan(ch)

	for num := range result {
		fmt.Printf("Square num: %d\n", num)
		time.Sleep(500 * time.Millisecond)
	}

}

func arrToChan(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer log.Println("func 1 finish")
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
		defer log.Println("func 2 finish")
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

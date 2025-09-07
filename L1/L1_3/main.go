package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	n := 5 // кол-во воркеров

	for id := 1; id <= n; id++ {
		go Worker(id, ch)
	}

	for i := 0; ; i++ {
		ch <- i
		time.Sleep(300 * time.Millisecond)
	}

}

func Worker(workerID int, ch chan int) {
	for elem := range ch {
		fmt.Printf("worker %d, elem: %d\n", workerID, elem)
	}
}

package main

import (
	"log"
	"sync"

	l118 "github.com/tmozzze/WB_techoschool/L1/L1_18"
)

func main() {
	counter := l118.NewConcurrentCounter()
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			counter.IncrementCounter()
			log.Printf("gorutine %d finished\n", i)
			wg.Done()
		}()
	}
	wg.Wait()
	counter.PrintCounter()

}

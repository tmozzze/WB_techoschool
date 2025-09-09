package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for elem := range ch {

			log.Println("Gorutine is working...")
			fmt.Printf("elem: %d\n", elem)
			time.Sleep(500 * time.Millisecond)

		}
		log.Println("Gorutine ended. Range channel stop.")
	}()

	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)

	wg.Wait()

}

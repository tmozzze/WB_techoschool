package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	timer := time.After(3 * time.Second)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-timer:
				log.Println("Gorutine ended. runtime.Goexit stop.")
				runtime.Goexit()
			default:
				log.Println("Gorutine is working...")
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()

}

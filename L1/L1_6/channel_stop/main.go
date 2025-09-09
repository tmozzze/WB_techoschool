package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopCh:
				log.Println("Gorutine ended. Channel stop.")
				return
			default:
				log.Println("Gorutine is working...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stopCh <- struct{}{}

	wg.Wait()

}

package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	stop := false
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if stop {
				log.Println("Gorutine ended. Condition stop.")
				return
			}
			log.Println("Gorutine is working...")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
	stop = true
	wg.Wait()

}

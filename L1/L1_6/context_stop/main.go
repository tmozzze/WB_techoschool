package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Println("Gorutine ended. Context stop.")
				return
			default:
				log.Println("Gorutine is working...")
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()

}

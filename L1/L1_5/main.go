package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	dataCh := make(chan int)
	var wg sync.WaitGroup

	// Количество секунд работы программы
	n := 5
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	wg.Add(2)
	go Writer(ctx, dataCh, &wg)
	go Reader(ctx, dataCh, &wg)
	wg.Wait()
	log.Println("main ended.")

}

func Writer(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			log.Println("Writer ended")
			close(ch)
			return
		case ch <- i:
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func Reader(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Println("Reader ended")
			return
		case elem, ok := <-ch:
			if !ok {
				log.Println("Reader: chanel closed")
				return
			}
			fmt.Printf("elem: %d\n", elem)
		}
	}
}

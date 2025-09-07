package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int)
	n := 5 // кол-во воркеров

	ctx, cancel := context.WithCancel(context.Background()) // контекст с отменой

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT) // ловим СИГИНТ

	for id := 1; id <= n; id++ {
		go Worker(ctx, id, ch)
	}

	go func() {
		<-sigChan
		log.Println("Завершение программы...")
		cancel()
	}()

	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			log.Println("main() завершён")
			close(ch)
			return
		case ch <- i:
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func Worker(ctx context.Context, workerID int, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d завершён\n", workerID)
			return
		case elem, ok := <-ch:
			if !ok {
				log.Printf("worker %d: канал закрыт", workerID)
				return
			}
			fmt.Printf("worker %d elem: %d\n", workerID, elem)
		}
	}
}

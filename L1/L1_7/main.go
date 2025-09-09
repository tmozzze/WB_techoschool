package main

import (
	"fmt"
	"sync"
)

type ConcurentMap struct {
	mu    sync.Mutex
	store map[int]int
}

func NewCMap() *ConcurentMap {
	return &ConcurentMap{store: make(map[int]int)}
}

func (cm *ConcurentMap) Set(key, val int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.store[key] = val
}

func Worker(workerID int, cm *ConcurentMap, wg *sync.WaitGroup) {
	for i := workerID*10 + 1; i <= workerID*10+3; i++ {
		cm.Set(i, i)
		fmt.Printf("Worker %d set key: %d, val: %d\n", workerID, i, i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	cMap := NewCMap()

	n := 3 // кол-во воркеров

	wg.Add(n)

	for id := 1; id <= n; id++ {
		go Worker(id, cMap, &wg)
	}

	wg.Wait()
	fmt.Println(cMap.store)

}

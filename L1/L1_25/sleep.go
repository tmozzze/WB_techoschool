package l125

import (
	"log"
	"time"
)

func Sleep(duration time.Duration) {
	log.Println("gorutine is sleeping...")
	defer log.Println("gorutine is waking up...")
	<-time.After(duration)
}

func CicleSleep(duration time.Duration) {
	start := time.Now()
	log.Println("gorutine is sleeping...")
	defer log.Println("gorutine is waking up...")
	for {
		if duration <= time.Since(start) {
			break
		}
	}
}

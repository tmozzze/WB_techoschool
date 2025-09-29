package main

import (
	"fmt"
	"time"

	l125 "github.com/tmozzze/WB_techoschool/L1/L1_25"
)

func main() {

	l125.Sleep(2 * time.Second)
	l125.CicleSleep(3 * time.Second)

	defer fmt.Println("main is ending...")
}

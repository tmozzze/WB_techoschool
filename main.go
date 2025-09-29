package main

import (
	"fmt"

	l124 "github.com/tmozzze/WB_techoschool/L1/L1_24"
)

func main() {

	p1 := l124.NewPoint(-1, 2)
	p2 := l124.NewPoint(2, -2)

	fmt.Println(p1.Distance(p2))

}

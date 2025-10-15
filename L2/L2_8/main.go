package main

import (
	"fmt"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/L2_8/ntp"
)

func main() {
	time, err := ntp.GetCurrentTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Current time : %s\n", time)
}

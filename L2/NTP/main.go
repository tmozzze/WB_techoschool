package main

import (
	"fmt"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/NTP/ntp"
)

func main() {
	time, err := ntp.GetCurrentTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Current time : %s\n", time)
}

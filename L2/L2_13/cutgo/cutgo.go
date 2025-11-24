package cutgo

import (
	"fmt"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/L2_13/cutgo_config"
)

func Cut() {
	// Init config
	cfg, err := cutgo_config.ParseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config error: %v", err)
		os.Exit(1)
	}

	cfg.Print()

}

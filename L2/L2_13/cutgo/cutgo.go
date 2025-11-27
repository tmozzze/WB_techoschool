package cutgo

import (
	"bufio"
	"fmt"
	"io"
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

	//TODO: Delete
	cfg.Print()

	// Init reader
	reader, cleanup, err := getInputReader(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reader error: %v", err)
		os.Exit(1)
	}
	defer cleanup()

	// Init writer
	writer := os.Stdout

	// Procces
	cut(reader, writer, cfg)
}

func cut(reader io.Reader, writer io.Writer, cfg *cutgo_config.Config) {
	sep := cfg.Delimiter

	scanner := bufio.NewScanner(reader)
	fields := cfg.Fields.GetSlice()

	for scanner.Scan() {
		fmt.Println(fields, sep)
	}
}

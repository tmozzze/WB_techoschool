package cutgo

import (
	"fmt"
	"io"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/L2_13/cutgo_config"
)

// getInputReader - get fileName from config choose file or STDIN reader
func getInputReader(cfg *cutgo_config.Config) (io.Reader, func(), error) {
	if cfg.FileName == "" {
		return os.Stdin, func() {}, nil
	}

	file, err := os.Open(cfg.FileName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}

	cleanup := func() {
		file.Close()
	}

	return file, cleanup, nil
}

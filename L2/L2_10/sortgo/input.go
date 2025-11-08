package sortgo

import (
	"fmt"
	"io"
	"os"
)

func getInputReader(args []string) (io.Reader, func(), error) {
	if len(args) == 0 {
		return os.Stdin, func() {}, nil
	}

	file, err := os.Open(args[0])
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file: %w", err)
	}

	cleanup := func() {
		file.Close()
	}

	return file, cleanup, nil
}

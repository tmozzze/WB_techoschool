package grepgo

import (
	"io"
	"os"
	"testing"

	"github.com/tmozzze/WB_techoschool/L2/L2_12/grepgo_config"
)

func TestGetInputReader(t *testing.T) {
	// 1. STDIN
	t.Run("Read from Stdin", func(t *testing.T) {
		cfg := &grepgo_config.Config{FileName: ""}

		reader, cleanup, err := getInputReader(cfg)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		defer cleanup()

		if reader != os.Stdin {
			t.Errorf("Expected os.Stdin, got %v", reader)
		}
	})

	// 2. File
	t.Run("Read from valid file", func(t *testing.T) {
		// Create file
		tmpFile, err := os.CreateTemp("", "test_grep_*.txt")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		// Delete file after test
		defer os.Remove(tmpFile.Name())

		// Fill file
		testContent := "Hello, World!"
		if _, err := tmpFile.WriteString(testContent); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		tmpFile.Close()

		// Init reader
		cfg := &grepgo_config.Config{FileName: tmpFile.Name()}
		reader, cleanup, err := getInputReader(cfg)
		if err != nil {
			t.Fatalf("Unexpected error opening file: %v", err)
		}
		defer cleanup()

		// Check
		data, err := io.ReadAll(reader)
		if err != nil {
			t.Fatalf("Failed to read from reader: %v", err)
		}

		if string(data) != testContent {
			t.Errorf("Expected content %q, got %q", testContent, string(data))
		}
	})

	// 3. Check non existent file
	t.Run("Read from invalid file", func(t *testing.T) {
		cfg := &grepgo_config.Config{FileName: "non_existent_file_12345.txt"}

		reader, cleanup, err := getInputReader(cfg)

		// Waiting error
		if err == nil {
			t.Error("Expected error for non-existent file, got nil")
		}

		// Reader must be nil
		if reader != nil {
			t.Errorf("Expected nil reader, got %v", reader)
		}

		// Check cleanup
		if cleanup != nil {
			cleanup()
		}
	})
}

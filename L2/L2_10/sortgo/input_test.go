package sortgo

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestGetInputReader_Stdin(t *testing.T) {
	r, cleanup, err := getInputReader([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer cleanup()

	if r != os.Stdin {
		t.Errorf("expected os.Stdin, got %T", r)
	}
}

func TestGetInputReader_File(t *testing.T) {
	contect := "hello world"
	tmpFile, err := os.CreateTemp("", "testfile.txt")
	if err != nil {
		t.Fatalf("failed to create remp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(contect); err != nil {
		t.Fatalf("failed to write: %v", err)
	}

	r, cleanup, err := getInputReader([]string{tmpFile.Name()})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer cleanup()

	buf := new(strings.Builder)
	if _, err := io.Copy(buf, r); err != nil {
		t.Fatalf("failed to read: %v", err)
	}

	if buf.String() != contect {
		t.Errorf("expected %q, got %q", contect, buf.String())
	}
}

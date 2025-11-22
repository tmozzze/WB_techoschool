package sortgo

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo_config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo_model"
)

// readLines - Take reader and make slice of *Line{}
// (when -u given make unique lines)
func readLines(reader io.Reader, flags *sortgo_config.Config) []*sortgo_model.Line {
	scanner := bufio.NewScanner(reader)
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	var lines []*sortgo_model.Line // slice for line obj
	for scanner.Scan() {
		raw := scanner.Text()

		// create new line object
		line := sortgo_model.NewLine(raw)
		// split fileds for line object
		line.SplitFields(flags.Sep)
		lines = append(lines, line)
	}

	return lines
}

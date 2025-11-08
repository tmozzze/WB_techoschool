package sortgo

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo/model"
)

func ReadLines(reader io.Reader, flags *config.Config) []*model.Line {
	scanner := bufio.NewScanner(reader)
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	var lines []*model.Line            // slice for line obj
	uniqueMap := make(map[string]bool) // map for check unique

	for scanner.Scan() {
		raw := scanner.Text()
		if flags.Unique && uniqueMap[raw] {
			continue
		}
		uniqueMap[raw] = true
		// create new line object
		line := model.NewLine(raw)
		// split fileds for line object
		line.SplitFields(flags.Sep)
		lines = append(lines, line)
	}

	return lines
}

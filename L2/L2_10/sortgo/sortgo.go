package sortgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/spf13/pflag"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
)

func printLines(lines []*Line, flags *config.Config) {
	// When -u given
	if flags.Reverse {
		for i := 0; i <= len(lines)-1; i++ {
			fmt.Println(lines[i].Raw)
		}
	} else {
		// Default
		for i := len(lines) - 1; i >= 0; i-- {
			fmt.Println(lines[i].Raw)
		}
	}
}

func readLines(reader io.Reader, flags *config.Config) []*Line {
	scanner := bufio.NewScanner(reader)
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	var lines []*Line                  // slice for line obj
	uniqueMap := make(map[string]bool) // map for check unique

	for scanner.Scan() {
		raw := scanner.Text()
		if flags.Unique && uniqueMap[raw] {
			continue
		}
		uniqueMap[raw] = true
		// create new line object
		line := NewLine(raw)
		// split fileds for line object
		line.splitFields(flags.Sep)
		lines = append(lines, line)
	}

	return lines
}

func getSortKey(line *Line, flags *config.Config) string {
	if flags.Key > 0 {
		return line.getField(flags.Key)
	}
	return line.Raw
}

func sortLines(lines []*Line, flags *config.Config) {
	sort.SliceStable(lines, func(i, j int) bool {
		keyI := getSortKey(lines[i], flags) // When -k given --> field
		keyJ := getSortKey(lines[j], flags) // When -k default --> raw

		// When -n given
		if flags.Num {
			numI, errI := strconv.ParseFloat(keyI, 64)
			numJ, errJ := strconv.ParseFloat(keyJ, 64)

			// When 1st and 2nd are not num
			if errI != nil && errJ != nil {
				return keyI < keyJ
			}
			// When 1st not num
			if errI != nil {
				return true
			}
			// When 2nd not num
			if errJ != nil {
				return false
			}

			if numI == numJ {
				return lines[i].Raw < lines[j].Raw
			}
			return numI < numJ
		}

		if keyI == keyJ {
			return lines[i].Raw < lines[j].Raw
		}

		return keyI < keyJ
	})
}

// Sort - make sorting string with flags
func Sort() {
	flags := config.ParseFlags()

	reader, cleanup, err := getInputReader(pflag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer cleanup()

	lines := readLines(reader, flags)

	sortLines(lines, flags)

	printLines(lines, flags)
}

// TODO: Refactoring code
//       Create any modes

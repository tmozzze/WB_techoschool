package sortgo

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/pflag"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/model"
)

// getSortKey - Get raw if -k = 0, else get field
func getSortKey(line *model.Line, flags *config.Config) string {
	var key string
	if flags.Key > 0 {
		key = line.GetField(flags.Key)
	} else {
		key = line.Raw
	}

	if flags.IgnoreTrailing {
		key = trimTrailingSpaces(key)
	}

	return key
}

// sortLines - make sorting slice of *Line
func sortLines(lines []*model.Line, flags *config.Config) {
	sort.SliceStable(lines, func(i, j int) bool {
		lineI := lines[i]
		lineJ := lines[j]

		switch {
		case flags.Month:
			// when -M given
			return monthSort(lineI, lineJ, flags)
		case flags.Num:
			// when -n given
			return numericSort(lineI, lineJ, flags)
		default:
			return stringSort(lineI, lineJ, flags)
		}

	})
}

// Sort - make sorting string with flags
func Sort() {
	// get flags
	flags := config.ParseFlags()

	reader, cleanup, err := getInputReader(pflag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer cleanup() // defer file.close()

	// make slie of *Line from text
	lines := readLines(reader, flags)

	// when -u given remove duplicates
	if flags.Unique {
		lines = removeDuplicates(lines, flags)
	}
	// sorting
	sortLines(lines, flags)
	// print results (when -r given make reverse)
	printLines(lines, flags)
}

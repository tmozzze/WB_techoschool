package sortgo

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/pflag"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo_config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo_model"
)

// getSortKey - Get raw if -k = 0, else get field
func getSortKey(line *sortgo_model.Line, flags *sortgo_config.Config) string {
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
func sortLines(lines []*sortgo_model.Line, flags *sortgo_config.Config) {
	sort.SliceStable(lines, func(i, j int) bool {
		lineI := lines[i]
		lineJ := lines[j]

		var result bool

		switch {
		case flags.Month:
			// when -M given
			result = monthSort(lineI, lineJ, flags)
		case flags.Num || flags.Human:
			// when -n or -h given
			result = numericSort(lineI, lineJ, flags)
		default:
			result = stringSort(lineI, lineJ, flags)
		}

		if flags.Reverse {
			return !result
		}

		return result

	})
}

// Sort - make sorting string with flags
func Sort() {
	// get flags
	flags := sortgo_config.ParseFlags()

	reader, cleanup, err := getInputReader(pflag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer cleanup() // defer file.close()

	// make slie of *Line from text
	lines := readLines(reader, flags)

	// when -c given
	if flags.Check {
		if !checkSorted(lines, flags) {
			os.Exit(1)
		}
		return
	}

	// when -u given remove duplicates
	if flags.Unique {
		lines = removeDuplicates(lines, flags)
	}
	// sorting
	sortLines(lines, flags)
	// print results (when -r given make reverse)
	printLines(lines)
}

package sortgo

import (
	"fmt"
	"os"
	"sort"

	"github.com/spf13/pflag"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo/comporators"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo/model"
)

func getSortKey(line *model.Line, flags *config.Config) string {
	if flags.Key > 0 {
		return line.GetField(flags.Key)
	}
	return line.Raw
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

	lines := ReadLines(reader, flags)

	var cmp comporators.Comporator
	switch {
	case flags.Month:
		cmp = comporators.MonthComporator(flags.Key)
	case flags.Num:
		cmp = comporators.NumericComporator(flags.Key)
	default:
		cmp = comporators.StringComporator(flags.Key)
	}

	cmp = comporators.StableTieComporator(cmp)
	if flags.Reverse {
		cmp = comporators.ReverseComporator(cmp)
	}

	sort.SliceStable(lines, func(i, j int) bool {
		return cmp(lines[i], lines[j])
	})

	PrintLines(lines)
}

// TODO: Refactoring code
//       Create any modes

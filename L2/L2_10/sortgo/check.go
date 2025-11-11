package sortgo

import (
	"fmt"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/model"
)

func checkSorted(lines []*model.Line, flags *config.Config) bool {
	for i := 1; i < len(lines); i++ {
		lineI := lines[i-1]
		lineJ := lines[i]

		var outOfOrder bool

		switch {
		case flags.Month:
			outOfOrder = monthSort(lineI, lineJ, flags)
		case flags.Num || flags.Human:
			outOfOrder = numericSort(lineI, lineJ, flags)
		default:
			outOfOrder = stringSort(lineI, lineJ, flags)
		}

		if flags.Reverse {
			outOfOrder = !outOfOrder
		}

		if !outOfOrder {
			fmt.Fprintf(os.Stderr, "disorder at line %d: %s\n", i+1, lineJ.Raw)
			return false
		}
	}

	return true
}

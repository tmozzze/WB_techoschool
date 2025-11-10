package sortgo

import (
	"fmt"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/model"
)

func checkSorted(lines []*model.Line, flags *config.Config) bool {
	for i := 1; i < len(lines); i++ {
		linePrev := lines[i-1]
		lineCurr := lines[i]

		var outOfOrder bool

		switch {
		case flags.Month:
			outOfOrder = monthSort(linePrev, lineCurr, flags)
		case flags.Num:
			outOfOrder = numericSort(linePrev, lineCurr, flags)
		default:
			outOfOrder = stringSort(linePrev, lineCurr, flags)
		}

		if flags.Reverse {
			outOfOrder = !outOfOrder
		}

		if outOfOrder {
			fmt.Fprintf(os.Stderr, "disorder at line %d: %s\n", i+1, lineCurr.Raw)
			return false
		}
	}

	return true
}

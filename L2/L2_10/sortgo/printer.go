package sortgo

import (
	"fmt"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo/model"
)

func PrintLines(lines []*model.Line) {
	// Default
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i].Raw)
	}
}

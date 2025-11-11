package sortgo

import (
	"fmt"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/model"
)

// printLines - print lines
// (when -r given reverse print)
func printLines(lines []*model.Line) {
	for i := 0; i <= len(lines)-1; i++ {
		fmt.Println(lines[i].Raw)
	}
}

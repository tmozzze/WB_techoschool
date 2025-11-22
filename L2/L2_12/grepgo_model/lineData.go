package grepgo_model

import (
	"fmt"
	"io"

	"github.com/tmozzze/WB_techoschool/L2/L2_12/grepgo_config"
)

// LineData - struct for line data
type LineData struct {
	Num  int
	Text string
}

// NewLineData - returns *LineData
func NewLineData(num int, text string) *LineData {
	return &LineData{Num: num, Text: text}
}

// PrintLine - formatted print LineData
func (l *LineData) PrintLine(w io.Writer, cfg *grepgo_config.Config, isContext bool) {
	// when -n
	if cfg.Number {
		// ':' sep for match
		// '-' sep for context
		sep := ":"
		if isContext {
			sep = "-"
		}
		// print text in format (15:text)
		fmt.Fprintf(w, "%d%s%s\n", l.Num, sep, l.Text)
	} else {
		fmt.Fprintln(w, l.Text)
	}
}

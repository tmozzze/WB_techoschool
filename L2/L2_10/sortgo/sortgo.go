package sortgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/pflag"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
)

type Line struct {
	Raw    string
	Fields []string
}

func NewLine(raw string) *Line {
	return &Line{Raw: raw}
}

func (l *Line) splitFields(sep string) {
	l.Fields = strings.Split(l.Raw, sep)
}

func (l *Line) getField(key int) string {
	if key <= len(l.Fields) {
		return l.Fields[key-1]
	}
	return ""
}

func printLines(lines []*Line, reverse bool) {
	switch reverse {
	case false:
		for i := 0; i <= len(lines)-1; i++ {
			fmt.Println(lines[i].Raw)
		}
	case true:
		for i := len(lines) - 1; i >= 0; i-- {
			fmt.Println(lines[i].Raw)
		}
	}
}

func Sort() {
	cfg := config.NewConfig()

	pflag.IntVarP(&cfg.Key, "key", "k", 0, "Number of column(required)")
	pflag.BoolVarP(&cfg.Num, "num", "n", false, "Num-sort")
	pflag.BoolVarP(&cfg.Reverse, "reverse", "r", false, "Reverse output")
	pflag.BoolVarP(&cfg.Unique, "unique", "u", false, "Unique output")
	pflag.Parse()
	cfg.Print()

	if cfg.Key < 0 {
		fmt.Fprintf(os.Stderr, "incorrect flag %v\n", cfg.Key)
		os.Exit(1)
	}

	var reader io.Reader = os.Stdin
	if args := pflag.Args(); len(args) > 0 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	scanner := bufio.NewScanner(reader)
	var lines []*Line
	for scanner.Scan() {
		raw := scanner.Text()
		// create new line object
		line := NewLine(raw)
		// split fileds for line object
		line.splitFields(cfg.Sep)
		lines = append(lines, line)
	}

	// When -k default
	if cfg.Key == 0 {
		sort.SliceStable(lines, func(i, j int) bool {
			return lines[i].Raw < lines[j].Raw
		})
	} else {
		// When -k given
		sort.SliceStable(lines, func(i, j int) bool {
			KeyI := lines[i].getField(cfg.Key)
			KeyJ := lines[j].getField(cfg.Key)

			// stable when fields equal
			if KeyI == KeyJ {
				return lines[i].Raw < lines[j].Raw
			}

			return KeyI < KeyJ
		})
	}

	printLines(lines, cfg.Reverse)
}

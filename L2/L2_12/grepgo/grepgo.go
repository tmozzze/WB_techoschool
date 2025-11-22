package grepgo

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/tmozzze/WB_techoschool/L2/L2_12/grepgo_config"
	"github.com/tmozzze/WB_techoschool/L2/L2_12/grepgo_model"
)

const (
	CONTEXT = true
	MATCH   = false
)

func Grep() {
	// init config
	cfg, err := grepgo_config.ParseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	// If you want print cfg
	// cfg.Print()

	// init reader
	reader, cleanup, err := getInputReader(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	// Init writer
	writer := os.Stdout

	// init matcher
	matcher, err := NewMatcher(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	// Run Check
	if err := searchPattern(writer, reader, matcher, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

}

func searchPattern(w io.Writer, r io.Reader, matcher Matcher, cfg *grepgo_config.Config) error {
	scanner := bufio.NewScanner(r)

	// Buffer for lines before matching (-B)
	beforeBuf := make([]*grepgo_model.LineData, 0, cfg.Before)

	matchCount := 0

	lastPrintedLine := -1
	linesAfterCounter := 0

	curLineNum := 0 // nums of lines starting from 1

	for scanner.Scan() {
		curLineNum++
		text := scanner.Text()

		lineObj := grepgo_model.NewLineData(curLineNum, text)

		isMatch := matcher(text)

		if isMatch {
			// When line matched
			matchCount++

			// When -c
			if cfg.Count {
				continue
			}

			// Print Before context
			for _, l := range beforeBuf {
				if l.Num > lastPrintedLine {
					l.PrintLine(w, cfg, CONTEXT)
					lastPrintedLine = l.Num
				}
			}
			beforeBuf = beforeBuf[:0] // Clean buf after printing

			// Print match line
			if lineObj.Num > lastPrintedLine {
				lineObj.PrintLine(w, cfg, MATCH)
			}

			linesAfterCounter = cfg.After

		} else {
			// NOT matched

			// Print After context
			if linesAfterCounter > 0 {
				if lineObj.Num > lastPrintedLine {
					lineObj.PrintLine(w, cfg, CONTEXT)
					lastPrintedLine = lineObj.Num
				}
				linesAfterCounter--
			} else {
				if cfg.Before > 0 {
					beforeBuf = append(beforeBuf, lineObj)

					// FIFO buffer
					if len(beforeBuf) > cfg.Before {
						beforeBuf = beforeBuf[1:]
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	if cfg.Count {
		fmt.Fprintln(w, matchCount)
	}

	return nil

}

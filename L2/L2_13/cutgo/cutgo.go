package cutgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tmozzze/WB_techoschool/L2/L2_13/cutgo_config"
)

func Cut() {
	// Init config
	cfg, err := cutgo_config.ParseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config error: %v", err)
		os.Exit(1)
	}

	// Init reader
	reader, cleanup, err := getInputReader(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reader error: %v", err)
		os.Exit(1)
	}
	defer cleanup()

	// Init writer
	writer := os.Stdout

	// Procces
	Cutgo(reader, writer, cfg)
}

func Cutgo(reader io.Reader, writer io.Writer, cfg *cutgo_config.Config) error {
	var text string

	sep := cfg.Delimiter

	scanner := bufio.NewScanner(reader)

	// sorting diapozone indexes
	cfg.Fields.SortFields()
	// get diapozone indexes
	diapozoneIdxes := cfg.Fields.GetSlice()

	for scanner.Scan() {
		text = scanner.Text()
		// check separator in line
		// If line HAS NOT separator
		if !strings.Contains(text, sep) {
			// and -s getted - Do nothing
			if cfg.Separated {
				continue
			} else {
				// and -s NOT getted - print full line
				fmt.Fprintln(writer, text)
				continue
			}
		}
		// split text on parts by separator
		fields := strings.Split(text, sep)

		var selectedFields []string // slice for fields in diapozone

		for _, idx := range diapozoneIdxes {
			// diapozone idx start from 1 (slice idx start from 0)
			formatedIdx := idx - 1

			if formatedIdx >= 0 && formatedIdx < len(fields) {
				selectedFields = append(selectedFields, fields[formatedIdx])
			}
		}
		// join slice in one line
		result := strings.Join(selectedFields, sep)
		fmt.Fprintln(writer, result)

	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	return nil
}

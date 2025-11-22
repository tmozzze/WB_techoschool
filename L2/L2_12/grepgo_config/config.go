package grepgo_config

import (
	"fmt"

	"github.com/spf13/pflag"
)

// Config - holds all command-line flags and settings.
type Config struct {
	After   int  // default 0 - means not used     (-A)
	Before  int  // default 0 - means not used     (-B)
	Context int  // default 0 - means not used     (-C)
	Count   bool // default false - means not used (-c)
	Ignore  bool // default false - means not used (-i)
	Invert  bool // default false - means not used (-v)
	Fixed   bool // default false - means not used (-F)
	Number  bool // default false - means not used (-n)

	Pattern  string
	FileName string // default "" - if STDIN
}

// NewConfig creates a new Config with default values.
func NewConfig() *Config {
	return &Config{
		After:   0,
		Before:  0,
		Context: 0,
		Count:   false,
		Ignore:  false,
		Invert:  false,
		Fixed:   false,
		Number:  false,
	}
}

// Print - prints the current configuration to stdout.
func (c *Config) Print() {
	fmt.Printf(
		"Config: -A: %v | -B: %v | -C: %v | -c: %v |\n"+
			" -i: %v | -v: %v | -F: %v | -n: %v\n",
		c.After,
		c.Before,
		c.Context,
		c.Count,
		c.Ignore,
		c.Invert,
		c.Fixed,
		c.Number,
	)
}

// ParseFlags - parse flags and config settings
func ParseFlags() (*Config, error) {
	cfg := NewConfig()

	pflag.IntVarP(&cfg.After, "after", "A", 0, "Print N lines after each line found")
	pflag.IntVarP(&cfg.Before, "before", "B", 0, "Print N lines before each line found")
	pflag.IntVarP(&cfg.Context, "context", "C", 0, "Print N lines before and after each line found")
	pflag.BoolVarP(&cfg.Count, "count", "c", false, "Print number of matched lines")
	pflag.BoolVarP(&cfg.Ignore, "ignore-case", "i", false, "Ignore case when searching")
	pflag.BoolVarP(&cfg.Invert, "invert-match", "v", false, "Print non-matched lines")
	pflag.BoolVarP(&cfg.Fixed, "fixed-string", "F", false, "Fixed string match")
	pflag.BoolVarP(&cfg.Number, "number", "n", false, "Print number of line")

	pflag.Parse()

	if cfg.After < 0 || cfg.Before < 0 || cfg.Context < 0 {
		return nil, fmt.Errorf("flag cannot be less than 0")
	}

	if cfg.Context > 0 {
		cfg.After = cfg.Context
		cfg.Before = cfg.Context
	}

	args := pflag.Args()
	if len(args) < 1 {
		return nil, fmt.Errorf("pattern is required")
	}

	cfg.Pattern = args[0]

	if len(args) > 1 {
		cfg.FileName = args[1]
	}

	return cfg, nil
}

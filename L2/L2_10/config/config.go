package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

// Config - holds all command-line flags and settings.
type Config struct {
	Key            int    // default 0 - means not used     (-k)
	Num            bool   // default false - means not used (-n)
	Reverse        bool   // default false - means not used (-r)
	Unique         bool   // default false - means not used (-u)
	Month          bool   // default false - means not used (-M)
	IgnoreTrailing bool   // default false - means not used (-b)
	Check          bool   // default false - means not used (-c)
	Human          bool   // default false - means not used (-h)
	Sep            string // default \t
}

// NewConfig creates a new Config with default values.
func NewConfig() *Config {
	return &Config{
		Key:            0,
		Num:            false,
		Reverse:        false,
		Unique:         false,
		Month:          false,
		IgnoreTrailing: false,
		Check:          false,
		Human:          false,
		Sep:            "\t",
	}
}

// Print - prints the current configuration to stdout.
func (c *Config) Print() {
	fmt.Printf(
		"Config: -k: %v | -n: %v | -r: %v | -u: %v |\n"+
			" -M: %v | -b: %v | -c: %v | -h: %v sep: %q\n",
		c.Key,
		c.Num,
		c.Reverse,
		c.Unique,
		c.Month,
		c.IgnoreTrailing,
		c.Check,
		c.Human,
		c.Sep,
	)
}

// ParseFlags - parses command-line flags and returns the resulting Config.
func ParseFlags() *Config {
	cfg := NewConfig()

	pflag.IntVarP(&cfg.Key, "key", "k", 0, "Number of column")
	pflag.BoolVarP(&cfg.Num, "num", "n", false, "Num-sort")
	pflag.BoolVarP(&cfg.Reverse, "reverse", "r", false, "Reverse output")
	pflag.BoolVarP(&cfg.Unique, "unique", "u", false, "Unique output")
	pflag.BoolVarP(&cfg.Month, "month", "M", false, "Sort by month name (Jan, Feb...)")
	pflag.BoolVarP(&cfg.IgnoreTrailing, "ignore-trailing", "b", false, "ignore trailing spaces")
	pflag.BoolVarP(&cfg.Check, "check", "c", false, "check sort")
	pflag.BoolVarP(&cfg.Human, "human-readable", "h", false, "human-readable sort")

	pflag.Parse()

	if cfg.Key < 0 {
		fmt.Fprintf(os.Stderr, "incorrect flag %v\n", cfg.Key)
		os.Exit(1)
	}

	//cfg.Print()
	return cfg
}

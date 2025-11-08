package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type Config struct {
	Key     int    // default 0 - means not used
	Num     bool   // defaul false - means not used
	Reverse bool   // defaul false - means not used
	Unique  bool   // defaul false - means not used
	Sep     string // defaul \t
}

func NewConfig() *Config {
	return &Config{Key: 0, Num: false, Reverse: false, Unique: false, Sep: "\t"}
}

func (c *Config) Print() {
	fmt.Printf("Config: -k: %v | -n: %v | -r: %v | -u: %v | sep: %q\n", c.Key, c.Num, c.Reverse, c.Unique, c.Sep)
}

func ParseFlags() *Config {
	cfg := NewConfig()

	pflag.IntVarP(&cfg.Key, "key", "k", 0, "Number of column(required)")
	pflag.BoolVarP(&cfg.Num, "num", "n", false, "Num-sort")
	pflag.BoolVarP(&cfg.Reverse, "reverse", "r", false, "Reverse output")
	pflag.BoolVarP(&cfg.Unique, "unique", "u", false, "Unique output")
	pflag.Parse()

	if cfg.Key < 0 {
		fmt.Fprintf(os.Stderr, "incorrect flag %v\n", cfg.Key)
		os.Exit(1)
	}

	cfg.Print()
	return cfg
}

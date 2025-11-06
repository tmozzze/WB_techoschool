package config

import "fmt"

type Config struct {
	Key     int
	Num     bool
	Reverse bool
	Unique  bool
	Sep     string
}

func NewConfig() *Config {
	return &Config{Key: 1, Num: false, Reverse: false, Unique: false, Sep: "\t"}
}

func (c *Config) Print() {
	fmt.Printf("Config: -k: %v | -n: %v | -r: %v | -u: %v | sep: %q\n", c.Key, c.Num, c.Reverse, c.Unique, c.Sep)
}

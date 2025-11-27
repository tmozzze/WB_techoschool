package cutgo_config

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/tmozzze/WB_techoschool/L2/L2_13/cutgo_model"
)

type Config struct {
	Fields    cutgo_model.IntDiapozoneValue // (-f) default len() == 0
	Delimiter string                        // sep - default '\t' (-d)
	Separated bool                          // default false (-s)

	FileName string // default "" - if STDIN
}

/*
-f "fields" — указание номеров полей (колонок),
которые нужно вывести. Номера через запятую, можно диапазоны.
Например: «-f 1,3-5» — вывести 1-й и с 3-го по 5-й столбцы.

-d "delimiter" — использовать другой разделитель (символ).
По умолчанию разделитель — табуляция ('\t').

-s – (separated) только строки, содержащие разделитель.
Если флаг указан, то строки без разделителя игнорируются (не выводятся).
*/

func NewConfig() *Config {
	return &Config{
		Fields:    make(cutgo_model.IntDiapozoneValue, 0),
		Delimiter: "\t",
		Separated: false,
	}
}

func (c *Config) Print() {
	fmt.Printf(
		"Config: -f: %v | -d: %q | -s: %v | file: %v |\n",
		c.Fields.String(),
		c.Delimiter,
		c.Separated,
		c.FileName,
	)
}

func ParseConfig() (*Config, error) {
	cfg := NewConfig()

	pflag.VarP(&cfg.Fields, "fields", "f", "number of printed column")
	pflag.StringVarP(&cfg.Delimiter, "delimiter", "d", "\t", "separator")
	pflag.BoolVarP(&cfg.Separated, "separated", "s", false, "print lines with separator/delimiter (default '\t')")

	pflag.Parse()

	if cfg.Fields.Len() == 0 {
		return nil, fmt.Errorf("-fields cannot be empty. Example (-f 1,3-5)")
	}

	args := pflag.Args()
	if len(args) > 1 {
		return nil, fmt.Errorf("too many arguments")
	}

	if len(args) == 1 {
		cfg.FileName = args[0]
	}

	return cfg, nil

}

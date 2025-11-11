package sortgo

import (
	"errors"
	"strconv"
	"strings"
)

var sizeSuffixes = map[string]float64{
	"k": 1024,
	"m": 1024 * 1024,
	"g": 1024 * 1024 * 1024,
	"t": 1024 * 1024 * 1024 * 1024,
	"p": 1024 * 1024 * 1024 * 1024 * 1024,
	"e": 1024 * 1024 * 1024 * 1024 * 1024 * 1024,
}

func parseHumanReadable(s string) (float64, error) {
	if s == "" {
		return 0, errors.New("empty string")
	}

	last := s[len(s)-1]
	multiplier := 1.0
	if factor, ok := sizeSuffixes[strings.ToLower(string(last))]; ok {
		multiplier = factor
		s = s[:len(s)-1]
	}

	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}

	return num * multiplier, nil
}

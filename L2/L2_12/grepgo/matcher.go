package grepgo

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/tmozzze/WB_techoschool/L2/L2_12/grepgo_config"
)

// Matcher - True when the line contains a Pattern
type Matcher func(line string) bool

// NewMathcer - returns Matcher
func NewMatcher(cfg *grepgo_config.Config) (Matcher, error) {
	pattern := cfg.Pattern

	// Fixed line and Ignore case
	if cfg.Fixed && cfg.Ignore {
		pattern = strings.ToLower(pattern)
	}

	var re *regexp.Regexp
	var err error
	if !cfg.Fixed {
		expr := pattern
		if cfg.Ignore {
			expr = "(?i)" + expr
		}
		re, err = regexp.Compile(expr)
		if err != nil {
			return nil, fmt.Errorf("failed to compile reg exp: %w", err)
		}
	}

	// Matcher
	matcher := func(line string) bool {
		matched := false

		if cfg.Fixed {
			// Fixed
			target := line

			// Ignore case
			if cfg.Ignore {
				target = strings.ToLower(line)
			}

			matched = strings.Contains(target, pattern)
		} else {
			// Regular
			matched = re.MatchString(line)
		}

		// Invert
		if cfg.Invert {
			return !matched
		}
		return matched
	}

	return matcher, nil
}

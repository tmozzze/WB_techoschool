package grepgo

import (
	"testing"

	"github.com/tmozzze/WB_techoschool/L2/L2_12/grepgo_config"
)

func TestMatcher(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		cfg     grepgo_config.Config
		line    string
		want    bool
	}{
		// Regexp
		{"Regexp simple", "abc", grepgo_config.Config{}, "xyz abc xyz", true},
		{"Regexp fail", "abc", grepgo_config.Config{}, "xyz ab xyz", false},
		{"Regexp dot", "a.c", grepgo_config.Config{}, "abc", true},
		{"Regexp ignore case", "ABC", grepgo_config.Config{Ignore: true}, "abc", true},

		// Fixed
		{"Fixed simple", "a.c", grepgo_config.Config{Fixed: true}, "a.c", true},
		{"Fixed dot literal", "a.c", grepgo_config.Config{Fixed: true}, "abc", false},
		{"Fixed ignore case", "ABC", grepgo_config.Config{Fixed: true, Ignore: true}, "abc", true},

		// Invert
		{"Invert true", "abc", grepgo_config.Config{Invert: true}, "xyz", true},
		{"Invert false", "abc", grepgo_config.Config{Invert: true}, "abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.cfg.Pattern = tt.pattern

			matcher, err := NewMatcher(&tt.cfg)
			if err != nil {
				t.Fatalf("NewMatcher error: %v", err)
			}

			got := matcher(tt.line)
			if got != tt.want {
				t.Errorf("Matcher(%q) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}

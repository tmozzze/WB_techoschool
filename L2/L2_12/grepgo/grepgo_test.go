package grepgo

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tmozzze/WB_techoschool/L2/L2_12/grepgo_config"
)

func TestGrepIntegration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		pattern  string
		cfg      grepgo_config.Config
		expected string
	}{
		// 1. Without flags
		{
			name:     "Without flags",
			input:    "apple\nbanana\ncherry",
			pattern:  "banana",
			cfg:      grepgo_config.Config{Pattern: "banana"},
			expected: "banana\n",
		},
		// 2. Ignore case (-i)
		{
			name:     "Ignore case (-i)",
			input:    "Apple\nBANANA\ncherry",
			pattern:  "banana",
			cfg:      grepgo_config.Config{Pattern: "banana", Ignore: true},
			expected: "BANANA\n",
		},
		// 3. Invert (-v)
		{
			name:     "Invert (-v)",
			input:    "A\nB\nC",
			pattern:  "B",
			cfg:      grepgo_config.Config{Pattern: "B", Invert: true},
			expected: "A\nC\n",
		},
		// 4. Fixed (-F)
		{
			name:     "Fixed (-F)",
			input:    "a.b\naXb",
			pattern:  "a.b",
			cfg:      grepgo_config.Config{Pattern: "a.b", Fixed: true},
			expected: "a.b\n",
		},
		// 5. Line numbers (-n)
		{
			name:     "Line numbers (-n)",
			input:    "one\ntwo\nthree",
			pattern:  "two",
			cfg:      grepgo_config.Config{Pattern: "two", Number: true},
			expected: "2:two\n",
		},
		// 6. Count (-c)
		{
			name:     "Count (-c)",
			input:    "match\nno\nmatch\nmatch",
			pattern:  "match",
			cfg:      grepgo_config.Config{Pattern: "match", Count: true},
			expected: "3\n",
		},
		// 7. After (-A)
		{
			name:     "After (-A)",
			input:    "1\n2 match\n3 after\n4",
			pattern:  "match",
			cfg:      grepgo_config.Config{Pattern: "match", After: 1},
			expected: "2 match\n3 after\n",
		},
		// 8. Before (-B)
		{
			name:     "Before (-B)",
			input:    "1 before\n2 match\n3",
			pattern:  "match",
			cfg:      grepgo_config.Config{Pattern: "match", Before: 1},
			expected: "1 before\n2 match\n",
		},
		// 9. Context (-C)
		{
			name:     "Context (-C)",
			input:    "1 before\n2 match\n3 after\n4",
			pattern:  "match",
			cfg:      grepgo_config.Config{Pattern: "match", Before: 1, After: 1},
			expected: "1 before\n2 match\n3 after\n",
		},
		// 10. Cross Context
		{
			name:     "Cross Context",
			input:    "1\n2 match\n3\n4 match\n5",
			pattern:  "match",
			cfg:      grepgo_config.Config{Pattern: "match", Before: 1, After: 1},
			expected: "1\n2 match\n3\n4 match\n5\n",
		},
		// 11. Number and Context (-nC)
		{
			name:     "Number and Context (-nC)",
			input:    "line1\nline2 match\nline3",
			pattern:  "match",
			cfg:      grepgo_config.Config{Pattern: "match", Before: 1, After: 1, Number: true},
			expected: "1-line1\n2:line2 match\n3-line3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			var output bytes.Buffer // writer

			// Init matcher
			matcher, err := NewMatcher(&tt.cfg)
			if err != nil {
				t.Fatalf("failed to create matcher: %v", err)
			}

			// Run
			if err := processInput(&output, reader, matcher, &tt.cfg); err != nil {
				t.Fatalf("failed to scan text: %v", err)
			}

			// Check
			got := output.String()
			if got != tt.expected {
				t.Errorf("\nInput:\n%s\nExpected:\n%q\nGot:\n%q", tt.input, tt.expected, got)
			}
		})
	}
}

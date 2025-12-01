package cutgo_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tmozzze/WB_techoschool/L2/L2_13/cutgo"
	"github.com/tmozzze/WB_techoschool/L2/L2_13/cutgo_config"
	"github.com/tmozzze/WB_techoschool/L2/L2_13/cutgo_model"
)

func TestCutIntegration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		cfg      *cutgo_config.Config
		expected string
	}{
		// 1. Basic field SEP - ':'
		{
			name:  "Basic fields",
			input: "root:x:0:0\ndaemon:x:1:1",
			cfg: &cutgo_config.Config{
				Delimiter: ":",
				Fields:    cutgo_model.IntDiapozoneValue{1, 3}, // -f 1,3
				Separated: false,
			},
			expected: "root:0\ndaemon:1\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			var writer bytes.Buffer

			err := cutgo.Cutgo(reader, &writer, tt.cfg)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if got := writer.String(); got != tt.expected {
				t.Errorf("\nInput:\n%q\nExpected:\n%q\nGot:\n%q", tt.input, tt.expected, got)

			}
		})
	}
}

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
			name:  "Basic sep = ':'",
			input: "1:2:0:3\n1:2:1:4",
			cfg: &cutgo_config.Config{
				Delimiter: ":",
				Fields:    cutgo_model.IntDiapozoneValue{1, 4}, // -f 1,3
				Separated: false,
			},
			expected: "1:3\n1:4\n",
		},

		// 2. Basic default field SEP - '\t'
		{
			name:  "Basic sep = '\t'",
			input: "1\t2\t0\t3\n1\t2\t1\t4",
			cfg: &cutgo_config.Config{
				Delimiter: "\t",
				Fields:    cutgo_model.IntDiapozoneValue{1, 4}, // -f 1,3
				Separated: false,
			},
			expected: "1\t3\n1\t4\n",
		},

		// 3. Without sep (-s = false)
		{
			name:  "Line hasn't delimiter, no -s flag",
			input: "hello world\nhello world",
			cfg: &cutgo_config.Config{
				Delimiter: ":",
				Fields:    cutgo_model.IntDiapozoneValue{1},
				Separated: false,
			},
			expected: "hello world\nhello world\n",
		},

		// 4. Without sep (-s = true)
		{
			name:  "Line hasn't delimiter, with -s flag",
			input: "hello world\nhello world",
			cfg: &cutgo_config.Config{
				Delimiter: ":",
				Fields:    cutgo_model.IntDiapozoneValue{1},
				Separated: true,
			},
			expected: "",
		},

		// 5. Line has mixed content (-s = true) (sep = ':')
		{
			name:  "Mixed content with -s",
			input: "hello1:world1\nhello2\tworld2\nhello3:world3",
			cfg: &cutgo_config.Config{
				Delimiter: ":",
				Fields:    cutgo_model.IntDiapozoneValue{1},
				Separated: true,
			},
			expected: "hello1\nhello3\n",
		},

		// 6. Line hasn't field (out of range)
		// 'hello:world' ask field = 5
		{
			name:  "Field out of range",
			input: "hello:world",
			cfg: &cutgo_config.Config{
				Delimiter: ":",
				Fields:    cutgo_model.IntDiapozoneValue{5},
				Separated: false,
			},
			expected: "\n",
		},

		// 7. Complex example
		// 'banana:3 \n apple\t2 \n cherry: 5:6'
		// flags: -d : -f 1-5 (1,2,3,4,5)
		{
			name:  "Complex example",
			input: "banana:3\napple\t2\ncherry: 5:6",
			cfg: &cutgo_config.Config{
				Delimiter: ":",
				Fields:    cutgo_model.IntDiapozoneValue{1, 2, 3, 4, 5},
				Separated: false,
			},
			expected: "banana:3\napple\t2\ncherry: 5:6\n",
		},

		// 8. DESC order fields and duplicates
		// fields = 3, 1, 1
		{
			name:  "DESC order fields and duplicates",
			input: "1:2:3",
			cfg: &cutgo_config.Config{
				Delimiter: ":",
				Fields:    cutgo_model.IntDiapozoneValue{3, 1, 1}, // flag 3, 1
				Separated: false,
			},
			expected: "1:3\n",
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

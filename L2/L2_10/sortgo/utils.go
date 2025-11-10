package sortgo

import (
	"strings"
	"unicode"
)

// TrimTrailingSpaces - remove trailing spaces from string
func trimTrailingSpaces(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

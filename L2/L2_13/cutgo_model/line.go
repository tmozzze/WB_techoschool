package cutgo_model

import "strings"

// Line - represents a single line with raw text and its fields
type Line struct {
	Text      string
	Fields    []string
	Separated bool
}

// NewLine - creates Line and returns *Line
func NewLine(text string) *Line {
	return &Line{Text: text}
}

// SplitFields - splits text by separator
func (l *Line) SplitFields(sep string) {
	l.Fields = strings.Split(l.Text, sep)
	if len(l.Fields) >= 2 {
		l.Separated = true
	}
}

// GetField - return Field by key(int) | "" if Field is empty
func (l *Line) GetField(key int) string {
	if key <= len(l.Fields) {
		return l.Fields[key]
	}
	return ""
}

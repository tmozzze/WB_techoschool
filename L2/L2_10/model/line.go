package model

import "strings"

// Line - represents a single line with raw text and its fields.
type Line struct {
	Raw    string
	Fields []string
}

// NewLine - creates a new Line with the given raw text.
func NewLine(raw string) *Line {
	return &Line{Raw: raw}
}

// SplitFields - take the Raw and split by separator
func (l *Line) SplitFields(sep string) {
	l.Fields = strings.Split(l.Raw, sep)
}

// GetField - returns the field at the given key (1-based). Returns "" if key is out of range.
func (l *Line) GetField(key int) string {
	if key <= len(l.Fields) {
		return l.Fields[key-1]
	}
	return ""
}

package model

import "strings"

type Line struct {
	Raw    string
	Fields []string
}

func NewLine(raw string) *Line {
	return &Line{Raw: raw}
}

// splitFields - take raw and split by separator
func (l *Line) SplitFields(sep string) {
	l.Fields = strings.Split(l.Raw, sep)
}

// getField - fill empty field
func (l *Line) GetField(key int) string {
	if key <= len(l.Fields) {
		return l.Fields[key-1]
	}
	return ""
}

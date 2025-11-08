package model

import (
	"strings"
)

type Line struct {
	Raw    string
	Fields []string
}

func NewLine(raw string) *Line {
	return &Line{Raw: raw}
}

func (l *Line) SplitFields(sep string) {
	l.Fields = strings.Split(l.Raw, sep)
}

func (l *Line) GetField(key int) string {
	if key <= 0 || key > len(l.Fields) {
		return l.Raw
	}
	return l.Fields[key-1]
}

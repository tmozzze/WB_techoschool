package sortgo

import "strings"

type Line struct {
	Raw    string
	Fields []string
}

func NewLine(raw string) *Line {
	return &Line{Raw: raw}
}

func (l *Line) splitFields(sep string) {
	l.Fields = strings.Split(l.Raw, sep)
}

func (l *Line) getField(key int) string {
	if key <= len(l.Fields) {
		return l.Fields[key-1]
	}
	return ""
}

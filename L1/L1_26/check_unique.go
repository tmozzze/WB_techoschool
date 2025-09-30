package l126

import "strings"

func ChechUnique(str string) bool {
	if len(str) <= 1 {
		return true
	}
	str = strings.ToLower(str)
	runes := []rune(str)
	m := make(map[rune]int)

	for _, elem := range runes {
		if m[elem] != 0 {
			return false
		}
		m[elem]++

	}
	return true
}

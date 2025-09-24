package l119

func ReverseString(str string) string {
	runeStr := []rune(str)

	if len(runeStr) < 2 {
		return str
	}

	left := 0
	right := len(runeStr) - 1

	for left < right {
		runeStr[left], runeStr[right] = runeStr[right], runeStr[left]
		left++
		right--
	}
	return string(runeStr)
}

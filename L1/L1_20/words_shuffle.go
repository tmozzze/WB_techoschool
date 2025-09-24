package l120

func ReverseRunes(runes []rune, start, end int) []rune {

	if len(runes) < 2 {
		return runes
	}

	left := start
	right := end

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return runes
}

func ShaffleWords(str string) string {
	runes := []rune(str)

	// reverse string
	revRunes := ReverseRunes(runes, 0, len(runes)-1)

	// reverse words
	start := 0
	for i := 0; i <= len(runes)-1; i++ {
		if revRunes[i] == ' ' {
			revRunes = ReverseRunes(revRunes, start, i-1)
			start = i + 1
		}
	}

	// reverse last word
	revRunes = ReverseRunes(revRunes, start, len(revRunes)-1)

	return string(revRunes)
}

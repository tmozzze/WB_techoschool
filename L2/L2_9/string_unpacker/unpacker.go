package stringunpacker

import (
	"fmt"
	"strconv"
	"unicode"
)

func Unpack(str string) (string, error) {
	runes := []rune(str)

	// Checking empty string
	if len(runes) == 0 {
		return "", nil
	}

	var result []rune
	escaped := false

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if escaped {
			result = append(result, r)
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			numStr := string(r)
			j := i + 1
			for j < len(runes) && unicode.IsDigit(runes[j]) {
				numStr += string(runes[j])
				j++
			}

			num, err := strconv.Atoi(numStr)
			if err != nil {
				return "", fmt.Errorf("invalid number: %v", err)
			}

			if num == 0 {
				return "", fmt.Errorf("zero num")
			}

			lastRune := result[len(result)-1]
			for k := 0; k < num-1; k++ {
				result = append(result, lastRune)
			}

			i = j - 1
			continue
		}

		result = append(result, r)
	}

	if escaped {
		return "", fmt.Errorf("alone backslash")
	}

	return string(result), nil

}

// g2

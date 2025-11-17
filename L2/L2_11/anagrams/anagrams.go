package anagrams

import (
	"sort"
	"strings"
)

// Make key for word
func formatKey(word string) string {
	word = strings.ToLower(word)
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func FindAnagrams(words []string) map[string][]string {
	if len(words) <= 1 {
		return nil
	}

	anagrams := make(map[string][]string)

	for _, word := range words {
		key := formatKey(word) // акпят - пятак пятка тяпка
		anagrams[key] = append(anagrams[key], strings.ToLower(word))

	}

	result := make(map[string][]string)

	for _, set := range anagrams {
		if len(set) <= 1 {
			continue
		}

		sort.Strings(set)
		result[set[0]] = set
	}

	return result
}

package interview_qs

import (
	"sort"
	"strings"
)

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)

	return strings.Join(chars, "")
}

func FindAnagrams(word string, wordList []string) []string {
	var anagrams []string

	sortedWord := sortString(strings.ToLower(word))

	for _, w := range wordList {
		if len(w) != len(word) {
			continue
		}

		if sortString(strings.ToLower(w)) == sortedWord {
			anagrams = append(anagrams, w)
		}
	}

	return anagrams
}

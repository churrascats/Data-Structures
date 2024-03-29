package trie

import "strings"

func ReturnLowerCaseAndTrimString(words ...string) []string {
	for index, word := range words {
		words[index] = strings.TrimSpace((strings.ToLower(word)))
	}

	return words
}

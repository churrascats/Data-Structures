package main

import (
	"data_structure/trie"
)

func main() {
	trie := trie.NewTrie()

	words := []string{
		"Carolina",
		"Rocambole",
		"Pastel",
		"Macaco",
	}

	for _, w := range words {
		trie.Insert(w)
	}

	trie.SearchAndPrint("Gilberto")

}

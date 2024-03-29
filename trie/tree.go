package trie

import "fmt"

const ALPHABET_SIZE = 26

type Node struct {
	Nodes [ALPHABET_SIZE]*Node
	IsEnd bool
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: &Node{},
	}
}

func (t *Trie) SearchAndPrint(word string) {
	word = ReturnLowerCaseAndTrimString([]string{word}...)[0]

	currentNode := t.Root
	for _, char := range word {
		charIndex := char - 'a'

		childrenNode := currentNode.Nodes[charIndex]
		if childrenNode == nil {
			childrenNode = &Node{}
		}
		currentNode = childrenNode
	}

	var result string
	if currentNode.IsEnd {
		result = fmt.Sprintf("Word %s found", word)
	} else {
		result = fmt.Sprintf("Word %s not found", word)
	}

	fmt.Println(result)
}

func (t *Trie) Insert(word string) {
	word = ReturnLowerCaseAndTrimString([]string{word}...)[0]

	currentNode := t.Root
	for _, char := range word {
		charIndex := char - 'a'
		childrenNode := &currentNode.Nodes[charIndex]
		if *childrenNode == nil {
			*childrenNode = &Node{}
		}
		currentNode = *childrenNode
	}
	currentNode.IsEnd = true
}

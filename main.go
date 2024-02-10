package convertwords

import (
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
	end      bool
	value    string
}

func New() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (t *TrieNode) Insert(word, value string) {
	node := t
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = New()
		}
		node = node.children[char]
	}
	node.end = true
	node.value = value
}

func (t *TrieNode) Search(s []rune) (string, int) {
	node := t
	for i, char := range s {
		if _, ok := node.children[char]; ok {
			node = node.children[char]
			if node.end {
				return node.value, i + 1
			}
		} else {
			break
		}
	}
	return "", 0
}

func ConvertWords(input string, root *TrieNode) string {
	var output strings.Builder
	runes := []rune(input)
	i := 0
	for i < len(runes) {
		word, length := root.Search(runes[i:])
		if length > 0 {
			output.WriteString(word)
			i += length
		} else {
			output.WriteRune(runes[i])
			i++
		}
	}
	return output.String()
}

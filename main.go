package main

import "fmt"

type TrieNode struct {
	word     string
	isWord   bool
	children []*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{"", false, make([]*TrieNode, 26)}
}

func add(node *TrieNode, word string) {
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = newTrieNode()
		}
		node = node.children[c-'a']
	}
	node.word = word
	node.isWord = true
}

func search_node(node *TrieNode, word string, i int) bool {
	if node == nil {
		return false
	}

	if i == len(word) {
		return node.isWord
	}

	c := word[i]

	if c == '.' {
		for j := 0; j < 26; j++ {

			if node.children[j] == nil {
				continue
			}

			search_result := search_node(node.children[j], word, i+1)
			if search_result {
				return true
			}
		}
	}

	return search_node(node.children[c-'a'], word, i+1)
}

func search(node *TrieNode, word string) bool {
	return search_node(node, word, 0)
}

func main() {
	trie := newTrieNode()

	word1 := "bad"
	word2 := "dad"
	word3 := "mad"

	add(trie, word1)
	add(trie, word2)
	add(trie, word3)

	search_word1 := "pad"
	search_word2 := "bad"
	search_word3 := ".ad"
	search_word4 := "b.."
	search_word5 := "b.d"

	fmt.Println("Search 'pad' = ", search(trie, search_word1))
	fmt.Println("Search 'bad' = ", search(trie, search_word2))
	fmt.Println("Search '.ad' = ", search(trie, search_word3))
	fmt.Println("Search 'b..' = ", search(trie, search_word4))
	fmt.Println("Search 'b.d' = ", search(trie, search_word5))
}

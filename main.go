package main

import "fmt"

type TrieNode struct {
	word     string
	isWord   bool
	children []*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		word:     "",
		isWord:   false,
		children: make([]*TrieNode, 26),
	}
}

func insert(node *TrieNode, word string) {
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = newTrieNode()
		}
		node = node.children[c-'a']
	}
	node.isWord = true
	node.word = word
}

func search(node *TrieNode, word string) bool {
	for _, c := range word {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return node.isWord
}

func startsWith(node *TrieNode, word string) bool {
	for _, c := range word {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return true
}

func main() {
	trie := newTrieNode()
	word1 := "apple"
	insert(trie, word1)
	fmt.Println("Search 'Apple' = ", search(trie, word1))
	word2 := "app"
	fmt.Println("Search 'App' = ", search(trie, word2))
	fmt.Println("Starts with 'App' = ", startsWith(trie, word2))
	insert(trie, word2)
	fmt.Println("Search 'App' = ", search(trie, word2))
}

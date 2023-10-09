package main

import "fmt"

type TrieNode struct {
	isWord   bool
	children []*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{false, make([]*TrieNode, 26)}
}

func add(node *TrieNode, word string) {
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = newTrieNode()
		}
		node = node.children[c-'a']
	}
	node.isWord = true
}

func search(node *TrieNode, board [][]rune, word string, i, j int, result *[]string) {
	m := len(board)
	n := len(board[0])

	if (i < 0) || (i >= m) || (j < 0) || (j >= n) || (board[i][j] == '#') {
		return
	}

	c := board[i][j]
	node = node.children[c-'a']

	if node == nil {
		return
	}

	word += string(c)

	if node.isWord {
		*result = append(*result, word)
		node.isWord = false
	}

	board[i][j] = '#'

	search(node, board, word, (i - 1), j, result)
	search(node, board, word, (i + 1), j, result)
	search(node, board, word, i, (j - 1), result)
	search(node, board, word, i, (j + 1), result)

	board[i][j] = c
}

func findWords(board [][]rune, words []string) []string {
	node := newTrieNode()
	result := make([]string, 0)

	for _, w := range words {
		add(node, w)
	}

	m := len(board)
	n := len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			search(node, board, "", i, j, &result)
		}
	}

	return result
}

func main() {
	board := [][]rune{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	words := []string{
		"oath",
		"pea",
		"eat",
		"rain",
	}

	result := findWords(board, words)

	fmt.Println("Result")
	for _, r := range result {
		fmt.Println("  -", r)
	}
}

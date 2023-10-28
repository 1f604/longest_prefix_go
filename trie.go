package main

import "strings"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children   map[rune]*TrieNode // runes are stored in the links
	isTerminal bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode), isTerminal: false}
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(s string) {
	node := t.root
	for _, r := range s {
		if _, ok := node.children[r]; !ok { // if value not in map
			node.children[r] = NewTrieNode()
		}
		node = node.children[r]
	}
	node.isTerminal = true
}

func (t *Trie) LongestMatchingPrefix(s string) string {
	node := t.root
	sb := strings.Builder{}
	result := ""
	for _, r := range s {
		if _, ok := node.children[r]; !ok { // if value not in map
			return result
		}
		sb.WriteRune(r)
		node = node.children[r]
		if node.isTerminal {
			result = sb.String()
		}
	}
	return result
}

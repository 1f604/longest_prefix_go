package main

import "strings"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode // runes are stored in the links
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
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
}

func (t *Trie) LongestMatchingPrefix(s string) string {
	node := t.root
	sb := strings.Builder{}
	for _, r := range s {
		if _, ok := node.children[r]; !ok { // if value not in map
			return sb.String()
		}
		sb.WriteRune(r)
		node = node.children[r]
	}
	return sb.String()
}

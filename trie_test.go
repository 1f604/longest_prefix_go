package main

import "testing"

// correctness check
func Test_Trie(t *testing.T) {
	prefixes := []string{
		"hello",
		"hello/world",
	}
	trie := NewTrie()
	for _, prefix := range prefixes {
		trie.Insert(prefix)
	}
	result := trie.LongestMatchingPrefix("hello world!")
	Assert_result_equals_bytes(t, []byte(result), nil, "hello", 1)

	result = trie.LongestMatchingPrefix("hell world!")
	Assert_result_equals_bytes(t, []byte(result), nil, "hell", 1)

	result = trie.LongestMatchingPrefix("hello/worldsdgasfg!")
	Assert_result_equals_bytes(t, []byte(result), nil, "hello/world", 1)
}

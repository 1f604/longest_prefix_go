package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

var letterBytes = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-")

func generate_n_level_prefix(level int, num int) []string {
	result := make([]string, 0, num)
	for i := 0; i < num; i++ {
		sb := strings.Builder{}
		sb.Grow(level * 8)

		for j := 0; j < level; j++ {
			for k := 0; k < 7; k++ {
				sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
			}
			sb.WriteByte('/')
		}
		result = append(result, sb.String())
	}
	return result
}

func setup(num_prefixes int, num_search_strings int) ([]string, []string) {
	// generate prefixes
	prefixes := []string{}
	for j := 0; j < 10; j++ { // generate up to 10 levels
		generated_prefixes := generate_n_level_prefix(j, num_prefixes)
		prefixes = append(prefixes, generated_prefixes...)
	}

	// generate search strings
	search_strings := []string{}
	for i := 0; i < num_search_strings; i++ {
		// find a random prefix
		random_prefix := prefixes[rand.Intn(len(prefixes))]
		sb := strings.Builder{}
		sb.WriteString(random_prefix)
		length := rand.Intn(15)
		for k := 0; k < length; k++ {
			sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
		}
		search_strings = append(search_strings, sb.String())
	}

	return prefixes, search_strings
}

var prefixes []string
var search_strings []string
var trie *Trie

func init() {
	prefixes, search_strings = setup(20, 1000)
	fmt.Println("num prefixes", len(prefixes))
	trie = NewTrie()
	for _, prefix := range prefixes {
		trie.Insert(prefix)
	}
}

func Benchmark_brute_force(b *testing.B) {
	//b.Fatal(len(prefixes))
	for _, search_string := range search_strings {
		BruteForce(prefixes, search_string)
	}
}

// correctness check
func Benchmark_Trie(b *testing.B) {
	for _, search_string := range search_strings {
		trie.LongestMatchingPrefix(search_string)
	}
}

// correctness check
func Test_brute_force(t *testing.T) {
	prefixes := []string{
		"hello",
		"hello/world",
	}
	result := BruteForce(prefixes, "hello world")
	Assert_result_equals_bytes(t, []byte(result), nil, "hello", 1)

	result = BruteForce(prefixes, "hello/world/hi")
	Assert_result_equals_bytes(t, []byte(result), nil, "hello/world", 1)
}

package main

import "strings"

func BruteForce(prefixes []string, input string) string {
	best := ""
	for _, prefix := range prefixes {
		if strings.HasPrefix(input, prefix) && len(best) < len(prefix) {
			best = prefix
		}
	}
	return best
}

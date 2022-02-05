package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pangrams("The quick brown fox jumps over the lazy dog"))
}

func pangrams(s string) string {
	// Write your code here
	ss := strings.Split(strings.ToLower(s), " ")
	s = strings.Join(ss, "")
	if len(s) < 26 {
		return "NOT PANGRAM"
	}
	m := map[rune]bool{}
	for _, v := range s {
		m[v] = true
	}
	if len(m) == 26 {
		return "PANGRAM"
	}
	return "NOT PANGRAM"
}

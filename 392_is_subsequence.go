package main

import (
	"strings"
)

func isSubsequence(s string, t string) bool {
	sChars := []rune(s)
	str := t
	for _, r := range sChars {
		i := strings.IndexRune(str, r)
		if i == -1 {
			return false
		}
		str = str[i+1:]
	}
	return true
}

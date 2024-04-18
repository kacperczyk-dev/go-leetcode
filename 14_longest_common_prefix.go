package main

import (
	"strings"
	"unicode/utf8"
)

func longestCommonPrefix(strs []string) string {
	shortest := strs[0]
	for _, v := range strs {
		if utf8.RuneCountInString(v) < utf8.RuneCountInString(shortest) {
			shortest = v
		}
	}
	for i := 0; i < len(strs); i++ {
		if shortest != "" && !strings.HasPrefix(strs[i], shortest) {
			runeSlice := []rune(shortest)
			trimmedString := runeSlice[:len(runeSlice)-1]
			shortest = string(trimmedString)
			i--
			continue
		}
	}
	return shortest
}

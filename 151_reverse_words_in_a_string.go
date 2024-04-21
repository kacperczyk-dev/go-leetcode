package main

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	arr := regexp.MustCompile("\\s{1,}").Split(s, -1)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	var sb strings.Builder
	for _, v := range arr {
		sb.WriteString(v)
		sb.WriteString(" ")
	}
	return strings.TrimSpace(sb.String())
}

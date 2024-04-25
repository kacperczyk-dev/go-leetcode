package main

import (
	"strings"
	"unicode/utf8"
)

func fullJustify(words []string, maxWidth int) []string {
	r := make([]string, 0)
	currCount := 0
	var sb strings.Builder
	for i := 0; i < len(words); {
		l := utf8.RuneCountInString(words[i])
		if currCount+l <= maxWidth {
			currCount += l
			sb.WriteString(words[i])
			if currCount < maxWidth {
				currCount += 1
				sb.WriteString(" ")
			}
			i++
		} else {
			r = append(r, strings.TrimSpace(sb.String()))
			sb.Reset()
			currCount = 0
		}
	}
	r = append(r, strings.TrimSpace(sb.String()))
	sb.Reset()
	for i, str := range r {
		wrds := strings.Split(str, " ")
		if len(wrds) == 1 {
			spaces := maxWidth - utf8.RuneCountInString(wrds[0])
			sb.WriteString(wrds[0])
			sb.WriteString(strings.Repeat(" ", spaces))
			r[i] = sb.String()
			sb.Reset()
			continue
		}
		if i == len(r)-1 {
			spaces := maxWidth - utf8.RuneCountInString(str)
			sb.WriteString(str)
			sb.WriteString(strings.Repeat(" ", spaces))
			r[i] = sb.String()
			sb.Reset()
			continue
		}
		l := 0
		for _, w := range wrds {
			l += utf8.RuneCountInString(w)
		}
		spaces := maxWidth - l
		for i := 0; spaces != 0; i++ {
			if i >= len(wrds)-1 {
				i = 0
			}
			sb.WriteString(wrds[i])
			sb.WriteString(" ")
			wrds[i] = sb.String()
			sb.Reset()
			spaces -= 1
		}
		r[i] = strings.Join(wrds, "")
	}
	return r
}

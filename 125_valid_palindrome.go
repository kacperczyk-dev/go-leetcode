package main

import "unicode"

func isPalindrome(s string) bool {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; {
		if !isAlphanumeric(chars[i]) {
			i++
			continue
		}
		if !isAlphanumeric(chars[j]) {
			j--
			continue
		}
		if unicode.ToLower(chars[i]) != unicode.ToLower(chars[j]) {
			return false
		}
		i++
		j--

	}
	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

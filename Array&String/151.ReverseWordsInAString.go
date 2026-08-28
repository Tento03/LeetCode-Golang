package main

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)

	left := 0
	right := len(words) - 1

	for left < right {
		words[left], words[right] = words[right], words[left]

		left++
		right--
	}

	return strings.Join(words, " ")
}

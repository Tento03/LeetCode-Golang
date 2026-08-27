package main

import "strings"

func reverseVowels(s string) string {
	arr := []rune(s)
	vowels := "aiueoAIUEO"

	left := 0
	right := len(arr) - 1

	for left < right {
		if !strings.ContainsRune(vowels, arr[left]) {
			left++
			continue
		}

		if !strings.ContainsRune(vowels, arr[right]) {
			right--
			continue
		}

		arr[left], arr[right] = arr[right], arr[left]

		left++
		right--
	}

	return string(arr)
}

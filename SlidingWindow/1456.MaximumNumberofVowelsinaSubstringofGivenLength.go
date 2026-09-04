package main

import "strings"

func maxVowels(s string, k int) int {
	sum := 0
	vowels := "aiueo"
	words := []rune(s)
	maxSum := 0

	for i := 0; i < k; i++ {
		if strings.ContainsRune(vowels, words[i]) {
			sum++
		}
	}

	maxSum = sum

	for i := k; i < len(words); i++ {
		if strings.ContainsRune(vowels, words[i]) {
			sum++
		}

		if strings.ContainsRune(vowels, words[i-k]) {
			sum--
		}

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

package main

import "strings"

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	n := gcd(len(str1), len(str2))

	candidate := str1[:n]

	if strings.Repeat(candidate, len(str1)/n) != str1 {
		return ""
	}

	if strings.Repeat(candidate, len(str2)/n) != str2 {
		return ""
	}

	return candidate
}

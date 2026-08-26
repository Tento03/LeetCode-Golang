package main

func mergeAlternately(word1 string, word2 string) string {
	n := len(word1)
	result := ""

	if n < len(word2) {
		n = len(word2)
	}

	for i := 0; i < n; i++ {
		if i < len(word1) {
			result += string(word1[i])
		}

		if i < len(word2) {
			result += string(word2[i])
		}
	}

	return result
}

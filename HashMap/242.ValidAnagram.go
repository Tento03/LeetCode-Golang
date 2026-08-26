package main

func isAnagram(s string, t string) bool {
	seen := map[rune]int{}

	for _, cha := range s {
		seen[cha]++
	}

	for _, cha := range t {
		seen[cha]--
	}

	for _, count := range seen {
		if count != 0 {
			return false
		}
	}

	return true
}

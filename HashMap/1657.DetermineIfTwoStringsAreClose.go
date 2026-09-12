package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	set1 := map[rune]int{}
	set2 := map[rune]int{}

	for _, word := range word1 {
		set1[word]++
	}

	for _, word := range word2 {
		set2[word]++
	}

	for word := range set1 {
		if _, ok := set2[word]; !ok {
			return false
		}
	}

	result1 := []int{}
	result2 := []int{}

	for _, num := range set1 {
		result1 = append(result1, num)
	}

	for _, num := range set2 {
		result2 = append(result2, num)
	}

	freq1 := map[int]int{}
	freq2 := map[int]int{}

	for _, num := range result1 {
		freq1[num]++
	}

	for _, num := range result2 {
		freq2[num]++
	}

	for num, count := range freq1 {
		if freq2[num] != count {
			return false
		}
	}

	return true
}

package main

func uniqueOccurrences(arr []int) bool {
	set := map[int]int{}

	for _, num := range arr {
		set[num]++
	}

	result := []int{}

	for _, num := range set {
		result = append(result, num)
	}

	seen := map[int]bool{}

	for _, num := range result {
		if seen[num] {
			return false
		}

		seen[num] = true
	}

	return true
}

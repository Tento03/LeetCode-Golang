package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxValue := candies[0]
	result := []bool{}

	for i := 1; i < len(candies); i++ {
		if candies[i] > maxValue {
			maxValue = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxValue {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}

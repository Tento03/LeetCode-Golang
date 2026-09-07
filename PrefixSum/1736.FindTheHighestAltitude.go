package main

func largestAltitude(gain []int) int {
	maxAltitude := 0
	sum := 0

	for i := 0; i < len(gain); i++ {
		sum += gain[i]

		if sum > maxAltitude {
			maxAltitude = sum
		}
	}

	return maxAltitude
}

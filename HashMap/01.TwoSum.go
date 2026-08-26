package main

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		value, ok := seen[target-num]

		if ok {
			return []int{value, i}
		}

		seen[num] = i
	}

	return []int{}
}

package main

func longestSubarray(nums []int) int {
	left := 0
	zeros := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}

			left++
		}

		windowLen := right - left + 1

		if windowLen-1 > maxLen {
			maxLen = windowLen - 1
		}
	}

	return maxLen
}

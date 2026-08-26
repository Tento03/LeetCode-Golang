package main

func maxSubArray(nums []int) int {
	var currValue = nums[0]
	var maxValue = nums[0]

	for i := 1; i < len(nums); i++ {
		currValue = max(currValue+nums[i], nums[i])

		if currValue > maxValue {
			maxValue = currValue
		}
	}

	return maxValue
}

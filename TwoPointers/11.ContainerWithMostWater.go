package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0

	for left < right {
		var shorter int

		if height[left] < height[right] {
			shorter = height[left]
		} else {
			shorter = height[right]
		}

		area := (right - left) * shorter

		if area > max {
			max = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

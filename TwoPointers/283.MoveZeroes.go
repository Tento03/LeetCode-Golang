package main

import "fmt"

func moveZeroes(nums []int) {
	pos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}

	for i := pos; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}

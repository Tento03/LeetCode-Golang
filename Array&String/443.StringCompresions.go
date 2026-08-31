package main

import "strconv"

func compress(chars []byte) int {
	read := 0
	write := 0

	for read < len(chars) {
		current := chars[read]
		count := 0

		for read < len(chars) && chars[read] == current {
			count++
			read++
		}

		chars[write] = current
		write++

		if count > 1 {
			for _, digit := range strconv.Itoa(count) {
				chars[write] = byte(digit)
				write++
			}
		}
	}

	return write
}

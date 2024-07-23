package main

import (
	"fmt"
)


func main() {
	numbers := []int{1, 2, 3, 4, 4, 9, 56, 90}
	fmt.Print(twoSum2(numbers, 8))
}

func twoSum2(numbers []int, target int) []int {
	// binary search to find complement
	for i := 0; i < len(numbers); i++ {
		num := numbers[i]
		comp := target - num

		_ = comp

		low, high := i+1, len(numbers)-1

		var count int
		for low <= high {
			mid := (low + high) / 2

			if comp == numbers[mid] {
				return []int{i + 1, mid + 1}
			} else if comp < numbers[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
			count++
			if count > 8 {
				break
			}
		}

	}
	return []int{}
}

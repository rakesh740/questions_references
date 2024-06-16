package main

import "fmt"

func main() {

	nums := []int{5}

	fmt.Println("Hello, 世界", search(nums, 5))
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high + low) / 2

		fmt.Println(high, low, mid)

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1

		}
		if nums[mid] < target {
			low = mid + 1
		}

	}

	return -1
}

func searchBinary(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high + low) / 2

		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			high = mid - 1
		case nums[mid] < target:
			low = mid + 1
		}

	}
	return -1
}

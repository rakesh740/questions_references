package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{7, -13, -1, 1, -6, 14, 10, -2, 1, 9, 11, -10, 8, -10, 14, 13, -1, 4, -6, -3, -5, 3, 3, 12, -5, 11, 5, -6, -2, 0, -6, 12, 3, 0, -2, 12, -1, -7, -5, 8, 10, 13, 13, 3, 10, 12, -7, -6, -7, -5, -1, 3, 5, -13, -8, -15, 13, 13, -14, -12, -2, -5, -15, 8, 11, -1, 6, -13, -1, 8, 10, -14, -1, 0, -4, -6, -3, 5, -4, -2, 7, 10, 8, -3, 12, -14, -10, 3, 14, -9, -2, -11, -6, -9, 13, 12, -3, 4, 14, 3, -11, 2, 5, -5, -13, -14, -3, -8}
	fmt.Print(threeSum(numbers))
}
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		num1 := nums[i]

		for j := 0; j < len(nums)-1; j++ {

			num2 := nums[j]
			comp := 0 - num1 - num2
			low, high := j+1, len(nums)-1

			var count int
			for low <= high {
				mid := (low + high) / 2

				if comp == nums[mid] {

					res = append(res, []int{num1, num2, nums[mid]})
				} else if comp < nums[mid] {
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

	}
	return res
}

func twoSum(numbers []int, target int) []int {
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

func threeSum1(nums []int) [][]int {
	var res [][]int

	// three sum means two two sum

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {

				if nums[i]+nums[j]+nums[k] == 0 {
					doappend := true
					for _, v := range res {
						if IsIn(v, []int{nums[i], nums[j], nums[k]}) {
							doappend = false
						}
					}
					if doappend {
						res = append(res, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return res
}

func IsIn(old []int, new []int) bool {
	var mapNums = make(map[int]int)
	for _, v := range old {
		mapNums[v]++
	}

	for _, v := range new {
		mapNums[v]--
	}

	for _, v := range mapNums {
		if v != 0 {
			return false
		}
	}
	return true
}

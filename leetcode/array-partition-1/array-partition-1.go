package main

import "fmt"

func main() {

	nums := []int{1, 8, 0, 9, 4, 5}

	fmt.Print(arrayPairSum(nums))
}

func mergeSort(nums []int) []int {

	n, mid := len(nums), len(nums)/2

	if n < 2 {
		return nums
	}

	var left, right []int

	for i := 0; i < mid; i++ {
		left = append(left, nums[i])
	}

	for i := mid; i < n; i++ {
		right = append(right, nums[i])
	}

	left = mergeSort(left)
	right = mergeSort(right)

	nums = merge(left, right)

	return nums
}

func merge(left, right []int) []int {

	// merge arr1, arr2 array

	var mergedArray []int

	var i, j int

	for i < len(left) && j < len(right) {
		if left[i] >= right[j] {
			mergedArray = append(mergedArray, left[i])
			i++
		} else {
			mergedArray = append(mergedArray, right[j])
			j++
		}

	}

	for i < len(left) {
		mergedArray = append(mergedArray, left[i])
		i++
	}

	for j < len(right) {
		mergedArray = append(mergedArray, right[j])
		j++
	}

	return mergedArray
}

func arrayPairSum(nums []int) int {

	sortedNums := mergeSort(nums)

	var lastIndex, sum int

	if len(sortedNums)%2 == 0 {
		lastIndex = len(sortedNums)
	} else {
		lastIndex = len(sortedNums) - 1
	}
	for i := 0; i < lastIndex; i += 2 {
		sum += sortedNums[i+1]
	}

	return sum
}

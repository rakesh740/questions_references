package main

import "fmt"

/*
 * Complete the compressString function below.
 */

func findKthLargest(nums []int, k int) int {
	nums = bubbleSort(nums)
	return nums[len(nums)-k]
}

func main() {

	fmt.Print()

}

func bubbleSort(input []int) []int {

	// n is the number of items in our list
	n := len(input)

	swapped := true

	for swapped {

		swapped = false

		for i := 1; i < n; i++ {

			if input[i-1] > input[i] {

				input[i], input[i-1] = input[i-1], input[i]

				swapped = true
			}
		}
	}
	// finally, print out the sorted list
	return input
}

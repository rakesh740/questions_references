package main

import "fmt"

/*
 * Complete the compressString function below.
 */

func getMaximumGenerated(n int) int {

	nums := make([]int, n+1)
	//var max int

	for i := 0; i <= n; i++ {

		if i == 0 {

			nums[i] = i
		} else if i == 1 {
			nums[i] = i
		} else if i%2 == 0 {
			p := i / 2
			nums[i] = nums[p]
		} else if i%2 != 0 {
			p := i / 2
			nums[i] = nums[p] + nums[p+1]
		}
	}

	nums = bubbleSort(nums)

	fmt.Print(nums)

	return nums[n]
}

func main() {

	fmt.Print(getMaximumGenerated(15))

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

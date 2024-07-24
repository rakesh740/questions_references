package main

import "fmt"

func main() {

	ar := []int{1, 2, 3, 4}
	pr := productExceptSelf(ar)
	fmt.Printf("pr: %v\n", pr)
}

func productExceptSelf(nums []int) []int {

	var output = make([]int, len(nums))
	var prefix int = 1
	var postfix int = 1

	for i := 0; i < len(output); i++ {
		output[i] = 1
	}

	for i := 0; i < len(nums); i++ {

		output[i] = output[i] * prefix

		prefix = prefix * nums[i]

	}

	fmt.Printf("output: %v\n", output)

	for i := len(output) - 1; i >= 0; i-- {
		output[i] = output[i] * postfix

		postfix = postfix * nums[i]
	}

	// multiply with every number then divide by each number as go through the loop

	return output
}

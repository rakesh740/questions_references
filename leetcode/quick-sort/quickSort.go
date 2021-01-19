package main

import "fmt"

func quickSort(arr []int, low, high int) []int {
	if low < high {
		/* pi is partitioning index, arr[pi] is now
		   at right place */
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)  // Before pi
		quickSort(arr, pi+1, high) // After pi
	}

	return arr
}

func partition(arr []int, low, high int) int {

	pivot := arr[high]

	i := (low - 1)

	for j := low; j <= high-1; j++ {

		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]

		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return (i + 1)
}

func main() {

	nums1 := []int{9, 8, 5, 3, 2}

	arr := quickSort(nums1, 0, 4)

	fmt.Print(arr)
}

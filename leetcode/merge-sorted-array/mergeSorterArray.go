package main

/*
 * Complete the compressString function below.
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	/*
	 * Write your code here.
	 */

	var newArray []int
	var i, j int

	i = 0
	j = 0

	for j < n && i < m {

		if nums1[i] <= nums2[j] {

			newArray = append(newArray, nums1[i])
			i++

		} else {

			newArray = append(newArray, nums2[j])
			j++

		}

	}

	if j <= n-1 {
		for j < n {
			newArray = append(newArray, nums2[j])
			j++
		}
	}

	if i <= m-1 {
		for i < m {
			newArray = append(newArray, nums1[i])
			i++
		}
	}

	for i = 0; i < m+n; i++ {
		nums1[i] = newArray[i]
	}

}

func main() {

	nums1 := []int{1}
	m, n := 1, 0
	nums2 := []int{}

	merge(nums1, m, nums2, n)

}

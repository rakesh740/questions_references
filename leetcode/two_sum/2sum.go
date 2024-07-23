package main

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		remain := target - nums[i]
		for j := i; j < len(nums); j++ {
			if remain == nums[j] {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}

	return result
}

func twoSumBetter(nums []int, target int) []int {

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if _, ok := m[complement]; ok {
			return []int{m[complement], i}
		}
		m[nums[i]] = i
	}
	return []int{}
}

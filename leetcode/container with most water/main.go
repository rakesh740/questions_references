package main

func main() {

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	curMax := 0
	for left < right {

		minHeight := min(height[left], height[right])

		curArea := minHeight * (right - left)
		curMax = max(curMax, curArea)

		switch minHeight {
		case height[left]:
			left++
		case height[right]:
			right--
		}
	}

	return curMax
}

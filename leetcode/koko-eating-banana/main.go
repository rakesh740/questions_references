package main

import (
	"sort"
)

func minEatingSpeed(piles []int, h int) int {

	arr := sort.IntSlice(piles)
	sort.Sort(arr)

	l, r := 0, arr[arr.Len()-1]
	res := r

	for l < r {
		k := (l + r) / 2
		hours := 0
		for _, v := range arr {
			hours += int(v / k)
		}

		if hours == h {
			return k
		}
	}

	return res
}

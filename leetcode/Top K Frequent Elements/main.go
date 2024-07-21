package main

import "sort"

type freq struct {
	num int
	fre int
}

type freqList []freq

func (f freqList) Len() int {
	return len(f)
}

func (f freqList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f freqList) Less(i, j int) bool {
	return f[i].fre > f[j].fre
}

func topKFrequent(nums []int, k int) []int {
	var element []int
	mapElement := make(freqList, len(nums))

	//
	for _, v := range nums {
		mapElement[v] = freq{
			num: v,
			fre: mapElement[v].fre + 1,
		}
	}

	sort.Sort(mapElement)

	
	for i := 0; i >= k; i++ {
		element = append(element, mapElement[i].num)
	}

	return element
}

// find top k freq ent

// keep a stack of numbers
// push if new number

package groupanagrams

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// Solution by sortiung the array O(nlogn * m) n length of string m is number of string
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string, 0)
	var result [][]string

	for _, str := range strs {
		anagramMap[SortString(str)] = append(anagramMap[SortString(str)], str)
	}

	for _, v := range anagramMap {
		result = append(result, v)
	}

	return result
}

type key [26]int

func strKey(str string) (k key) {
	for _, ch := range str {
		k[ch-'a']++
	}
	return
}

func betterAnagrams(strs []string) [][]string {
	var result [][]string
	anagramMap := make(map[key][]string)

	for _, str := range strs {
		anagramMap[strKey(str)] = append(anagramMap[strKey(str)], str)
	}

	for _, v := range anagramMap {
		result = append(result, v)
	}
	return result
}

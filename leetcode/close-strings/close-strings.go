package main

import "fmt"

//  Two strings are considered close if you can attain one from the other using the following operations:

// Operation 1: Swap any two existing characters.
// For example, abcde -> aecdb
// Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
// For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
// You can use the operations on either string as many times as necessary.

func main() {

	fmt.Print(closeStrings("hhhhhhhhhhh", "hhhhhhhhhhh"), "\n")
}

func closeStrings(word1 string, word2 string) bool {

	wordArr1 := []rune(word1)
	wordArr2 := []rune(word2)

	if diff(len(wordArr1), len(wordArr2)) > 0 {
		return false
	}

	table1 := make(map[rune]int)
	table2 := make(map[rune]int)

	unique := make(map[rune]int)

	for i := 0; i < len(wordArr1); i++ {
		table1[wordArr1[i]]++
		unique[wordArr1[i]]++
	}

	for i := 0; i < len(wordArr2); i++ {
		table2[wordArr2[i]]++

		unique[wordArr2[i]]--
		if unique[wordArr2[i]] == 0 {
			delete(unique, wordArr2[i])
		}
	}

	if diff(len(table1), len(table2)) > 0 {
		return false
	}

	if len(unique) > len(table1) {
		return false
	}

	table := make(map[int]int)

	for _, v1 := range table1 {

		table[v1]++
	}

	for _, v2 := range table2 {

		table[v2]--

	}

	for _, v := range table {
		if v != 0 {
			return false
		}
	}

	return true
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

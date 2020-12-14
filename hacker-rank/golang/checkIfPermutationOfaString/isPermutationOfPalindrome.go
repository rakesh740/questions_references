package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func buildCharFrequencyTable(phrase string) map[rune]int {
	letters := []rune(phrase)

	table := make(map[rune]int)

	for i := 0; i < len(letters); i++ {

		if table[letters[i]] >= 1 {
			table[letters[i]]--
			if table[letters[i]] == 0 {
				delete(table, letters[i])
			}
		} else {
			table[letters[i]]++
		}
	}
	return table
}

/*
 * Complete the isPermutationOfPalindrome function below.
 */
func isPermutationOfPalindrome(phrase string) bool {
	/*
	 * Write your code here.
	 */

	phrase = strings.ToLower(phrase)

	table := buildCharFrequencyTable(phrase)

	return checkMaxOneOdd(table)

}

func checkMaxOneOdd(table map[rune]int) bool {

	if len(table) > 1 {
		return false
	}
	return true
}

func main() {

	result := isPermutationOfPalindrome("Tcotcoook")

	fmt.Printf("%t \n", result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

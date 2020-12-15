package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func buildCharDifferenceTable(first, second string) map[rune]int {
	firstLetters := []rune(first)
	secondLetters := []rune(second)

	table := make(map[rune]int)

	for i := 0; i < len(firstLetters); i++ {
		table[firstLetters[i]]++
	}

	for i := 0; i < len(secondLetters); i++ {
		table[secondLetters[i]]--
		if table[secondLetters[i]] == 0 {
			delete(table, secondLetters[i])
		}
	}

	return table
}

/*
 * Complete the checkIfStringIsOneEditAway function below.
 */

func checkIfStringIsOneEditAway(first, second string) bool {
	/*
	 * Write your code here.
	 */

	first = strings.ToLower(first)
	second = strings.ToLower(second)

	table := buildCharDifferenceTable(first, second)

	return checkMaxOneOdd(table)
}

func checkMaxOneOdd(table map[rune]int) bool {

	for v := range table {
		fmt.Print(string(v), " ")
	}

	if len(table) > 2 {
		return false
	}
	return true
}

func main() {

	result := checkIfStringIsOneEditAway("palesp", "pales")

	fmt.Printf("\n%t \n", result)

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

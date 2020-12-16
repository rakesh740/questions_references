package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complete the compressString function below.
 */
func compressString(phrase string) string {
	/*
	 * Write your code here.
	 */

	letters := []rune(phrase)

	var newString []string
	var prevChar rune
	var count int

	for i := 0; i < len(letters); i++ {

		if i == 0 {
			prevChar = letters[i]
			count = 1
		} else {
			if i == len(letters)-1 {
				if letters[i] == prevChar {
					count++
					newString = append(newString, string(prevChar))
					newString = append(newString, strconv.Itoa(count))
				} else {
					newString = append(newString, string(prevChar))
					newString = append(newString, strconv.Itoa(count))
					prevChar = letters[i]
					count = 1
					newString = append(newString, string(prevChar))
					newString = append(newString, strconv.Itoa(count))
				}
			} else {
				if letters[i] == prevChar {
					count++
				} else {
					newString = append(newString, string(prevChar))
					newString = append(newString, strconv.Itoa(count))
					prevChar = letters[i]
					count = 1
				}
			}

		}
	}

	return strings.Join(newString, "")
}

func main() {

	result := compressString("aabcccccaaa")

	fmt.Printf("%s \n", result)

}

package main

import "fmt"

func main() {

	result := isPalindrome("abbbbbcbbbbba")

	fmt.Printf("%t \n", result)

}

func longestPalindrome(s string) string {

	return ""
}

func isPalindrome(s string) bool {

	letters := []rune(s)

	//fmt.Print(letters)

	for i := 0; i < (len(letters)+1)/2; i++ {

		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}

	return true
}

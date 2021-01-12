package main

import "fmt"

func isVowel(value rune) {

}

func main() {

	var T int
	fmt.Scanf("%d", &T)

	for i := 0; i < T; i++ {

		var myString string
		var count int

		fmt.Scanf("%s", &myString)

		myString = "nBBZLaosnm"

		for _, value := range myString {

			if value == 'a' || value == 'e' || value == 'i' || value == 'o' || value == 'u' {
				count++
			}

		}
		fmt.Printf("%d", count)
	}
}

package main

import "fmt"

// type groupChars []rune

// func isVowel(c rune) bool {
// 	vowels := groupChars{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
// 	for _, value := range vowels {
// 		if value == c {
// 			return true
// 		}
// 	}
// 	return false
// }

// func example() {
// 	myString := "OLapOKA3EOR"

//     fmt.Printf("%T\n", myString)

// 	// fmt.Printf("%T" , myString)
// 	t := 0
// 	for _, value := range myString {
// 		switch value {
// 		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
// 			t++
// 		}
// 	}
// 	fmt.Printf("%d Vowels.", t)

// }

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

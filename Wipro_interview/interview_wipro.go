package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func worker(ch, result chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for v := range ch {
		// do processing here
		fmt.Println(v, "is processed by ", i)
		result <- v
	}
}

func consumer(result chan int) {
	for v := range result {
		// do processing here
		fmt.Println(v)
	}
}

func main() {
	// ch := make(chan int)
	// result := make(chan int)
	// numWorker := 3
	// wg := sync.WaitGroup{}

	// wg.Add(1)
	// go produce(ch, &wg)

	// for i := 0; i < numWorker; i++ {
	// 	wg.Add(1)
	// 	go worker(ch, result, &wg, i)
	// }

	// go consumer(result)

	// wg.Wait()

	fmt.Println("Hello World! ",
		contains("at lest one Lowercase - at least one special char - at least one digit - len == 8", "len == 8"))

}

func produce(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

// print prime number between 1 and 100

// generate 5 random passwords - at least  one upperCase -
// at lest one Lowercase - at least one special char - at least one digit - len == 8

func passwordGenerator() {
	// regex -

	var passwords []string
	for i := 0; i < 5; i++ {
		passwords = append(passwords, generatePassword())
	}

	for _, v := range passwords {
		fmt.Printf("v: %v, len(v) %d \n", v, len(v))
	}

}

func generatePassword() string {

	// map [int]rune
	// get random number between 0, 7 - getDigit -
	// get random number between 0, 7 - getUpperCase - if already present find another rune + 1
	// get random number between 0, 7 - getLowerCase - if already present find another rune  pass
	// get random number between 0, 7 - getSpecialChar - if already present find another rune
	// [A d 0*  ]

	var funcArr []func() rune

	funcArr = append(funcArr, getDigit)
	funcArr = append(funcArr, getUpperCase)
	funcArr = append(funcArr, getLowerCase)
	funcArr = append(funcArr, getSpecialChar)

	var password []rune
	for i := 0; i < 8; i++ {

		password = append(password, getRandomFunc(funcArr)())
	}

	return string(password)
}

func getRandomFunc(funcArr []func() rune) func() rune {
	return funcArr[rand.Intn(len(funcArr))]
}

func getUpperCase() rune {
	i := rand.Intn(26)
	return 'A' + rune(i)
}

func getLowerCase() rune {
	i := rand.Intn(26)
	return 'a' + rune(i)
}

func getDigit() rune {
	return rune(rand.Intn(9))
}

func getSpecialChar() rune {
	return '&'
}

// search for string1 in string2 or not - string2 contain spaces - string1 contain spaces

//I am part of the string   -> []string 0, n

//am part -> []string 0 m

func contains(str1, str2 string) bool {
	n := len(str1)
	m := len(str2)
	for i := 0; i < n-m+1; i++ {
		j := i + m
		if checkEqual(str1[i:j], str2) {
			return true
		}
	}
	return false
}

func checkEqual(str1, str2 string) bool {
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

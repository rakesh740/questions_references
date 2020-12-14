package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func replaceSpaces(s string, truelength int) string {
	/*
	 * Write your code here.
	 */

	letters := []rune(s)

	var spaceCount, i int = 0, 0
	var index int

	for i = 0; i < truelength; i++ {
		if string(letters[i]) == " " {
			spaceCount++
		}
	}

	index = truelength + spaceCount*2

	for i = truelength - 1; i >= 0; i-- {
		if string(letters[i]) == " " {

			letters[index-1] = '0'
			letters[index-2] = '2'
			letters[index-3] = '%'
			index = index - 3

			fmt.Print(string(letters), "print here \n")

		} else {
			letters[index-1] = letters[i]
			fmt.Print(string(letters), "print \n")

			index--
		}
	}
	str := string(letters)

	return str
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// s := readLine(reader)

	result := replaceSpaces("Mr John Smith     ", 13)

	fmt.Printf("%s\n", result)

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

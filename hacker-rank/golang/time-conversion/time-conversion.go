package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	var arr2 []string
	var t string
	if strings.Contains(s, "PM") {
		arr2 = strings.Split(s, "P")
	}
	if strings.Contains(s, "AM") {

		arr2 = strings.Split(s, "A")
	}

	arr := strings.Split(arr2[0], ":")

	hh := arr[0]
	ht, _ := strconv.ParseInt(hh, 10, 10)

	var st string
	if strings.Contains(s, "PM") {

		if ht != 12 {
			ht = ht + 12

		}
		t = strconv.FormatInt(ht, 10)
	}
	if strings.Contains(s, "AM") {

		if ht == 12 {
			ht = 0
		}
		t = strconv.FormatInt(ht, 10)
	}

	if ht == ht%10 {
		t = "0" + t
	}
	arr[0] = t
	st = strings.Join(arr, ":")
	return st
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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

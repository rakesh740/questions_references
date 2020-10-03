package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {

	var pos, neg, zer float32

	for _, v := range arr {
		if v < 0 {
			neg++
		}
		if v > 0 {
			pos++
		}
		if v == 0 {
			zer++
		}
	}

	var posRatio, negRatio, zerRatio float32

	posRatio = pos / float32(len(arr))
	negRatio = neg / float32(len(arr))
	zerRatio = zer / float32(len(arr))

	fmt.Printf("%.6f\n", posRatio)
	fmt.Printf("%.6f\n", negRatio)
	fmt.Printf("%.6f\n", zerRatio)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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

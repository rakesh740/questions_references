package main

import (
	"fmt"
	"sync"
)

// given []string find all consecutive same chars in each word - find num of change needed to no consecutive char is same
//aaab -> 1
//abbb -> 1
//abaaaba -> 1
//abb -> 1

// given an array find arr[i] ^ arr[i+1] modulo (10^9 + 7) for each number

type cusError struct {
	msg string
}

func (c *cusError) Error() string {
	return c.msg
}

func f() error {
	var err *cusError
	return err
}

func main() {
	err := f()
	if err != nil {
		fmt.Println(err)
	}
}

type stru struct {
	mu  sync.Mutex
	arr []int
}

func create() {
	//p := new(stru)
	//print(p)
	//
	//var s stru
	//print(s)
}

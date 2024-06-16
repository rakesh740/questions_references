package main

import "fmt"

func main() {

	fmt.Println("Hello, 世界", isValid("(})"))
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	var stack []rune
	var head int = -1

	for i, v := range s {

		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
			head++
		case ')':
			if len(stack) == 0 || stack[head] != '(' {
				return false
			} else {
				stack = stack[:head]
				head--
			}

		case '}':
			if len(stack) == 0 || stack[head] != '{' {
				return false
			} else {
				stack = stack[:head]
				head--
			}
		case ']':
			if len(stack) == 0 || stack[head] != '[' {
				return false
			} else {
				stack = stack[:head]
				head--
			}
		}

		fmt.Println(stack, i)
		fmt.Println(head)
	}

	fmt.Println(head)

	return head == -1
}


// open with opening parenthesis
// can not open with closed paranthesis
// every parenthesis must be closed 
// every parenthesis must be closed in correct order 


func ALTisValid(s string) bool {
    if len(s) == 0 || len(s) % 2 == 1 {
        return false
    }
    
    pair := map[rune]rune {
        '(' : ')',
        '[' : ']',
        '{' : '}',
    }
    stack := []rune{}
    for _, r := range s {
        if _, ok := pair[r]; ok {
            stack = append(stack, r)
        } else if len(stack) == 0 || pair[stack[len(stack)-1]] != r {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
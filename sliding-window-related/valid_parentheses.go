package main

import "fmt"

/**
	algorithm : find needle in stack (one of sliding window)
	much faster
*/

/**
	s = haystack
	'(',')',{},() = needle
*/
func solve(s string) bool {
	stack := make([]rune, 0)
	
	for _, str := range s {
		if str == '(' || str == '{' || str == '[' {
			stack = append(stack, str)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				top := stack[len(stack)-1]
				if (str == ')' && top == '(') ||
				(str == '}' && top == '{') || 
				(str == ']' && top == '[') {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}

	return len(stack) == 0
}

func main() {
	input := "()[]{}"
	input2 := "(]"
	fmt.Println(solve(input))
	fmt.Println(solve(input2))
}
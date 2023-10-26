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
func needle_in_haystack_solve(s string) bool {
	// haystack
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

/*
	needle in haystack using map (sliding window)
	removing magic literal
*/
var closingToOpening = map[byte]byte {
    ')': '(',
    '}': '{',
    ']': '[',
}
func map_needle_in_haystack_solve(str string) bool {
	s := []byte(str)
    stack := make([]byte, 0)

    for _, ch := range s {
        matchingOpening, isClosing := closingToOpening[ch]

        if !isClosing { // it's opening
            stack = append(stack, ch)
            continue
        }

        if len(stack) == 0 { // closing, but no opening at all
            return false
        }

        lastOpening := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1] // "remove" it

        if lastOpening != matchingOpening {
            return false
        }
    }

    return len(stack) == 0
}

func main() {
	input := "()[]{}"
	input2 := "(]"
	fmt.Println(map_needle_in_haystack_solve(input))
	fmt.Println(map_needle_in_haystack_solve(input2))
}
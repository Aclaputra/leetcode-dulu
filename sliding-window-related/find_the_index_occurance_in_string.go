package main

import "fmt"

func sliding_window_simplified_solve(haystack string, needle string) int {
    for a, b := 0, len(needle); b <= len(haystack); a, b = a+1, b+1 {
        if haystack[a:b] == needle {
            return a
        }
    }
    return -1
}

func main() {
	input_haystack_1 := "sadbutsad"
	input_needle_1 := "sad"
	input_haystack_2 := "leetcode"
	input_needle_2 := "leeto"

	fmt.Println(sliding_window_simplified_solve(input_haystack_1, input_needle_1))
	fmt.Println(sliding_window_simplified_solve(input_haystack_2, input_needle_2))
}
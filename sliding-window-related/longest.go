package main

import (
	"fmt"
	"math"
)

/**
	disadvantage using char its only works 
	not that accurate
	still hardcode
	doesn't follow any algorithm convention
	should handle too much possible error like last index & first index
*/
func solve_backup(words string) int {
	var (
		char_set = [128]int{}
		topCount int
		count int
	)

	// if len(words) > 128 {
	// 	return -1
	// }
	for id, word := range words {
		// not unique anymore
		// fmt.Println(count, "count")
		// if id == 0 {
		// 	count++
		// }
		if char_set[word] > 0 {
			if count > topCount {
				topCount = count
			}
			count = 0
			char_set = [128]int{}
		}

		// still unique
		if char_set[word] <= 0 {
			char_set[word]++
			count++
		}

		// last index 
		if id == len(words)-1 {
			if count > topCount {
				topCount = count
			}
		}
		// fmt.Println(id, count, "count")
	}

	return topCount
}

// https://www.code-recipe.com/post/longest-substring-without-repeating-characters
func sliding_window_hashmap_01(s string) int {
	// Length of the given input string
	n := len(s)

	// Length of longest substring
	var result int

	// Map to store visited characters along with their index
	charIndexMap := make(map[uint8]int)

	// start indicates the start of current substring
	var start int

	// Iterate through the string and slide the window each time you find a duplicate
	// end indicates the end of current substring
	for end := 0; end < n; end++ {

		// Check if the current character is a duplicate
		duplicateIndex, isDuplicate := charIndexMap[s[end]]
		if isDuplicate {
			// Update the result for the substring in the current window before we found duplicate character
			result = int(math.Max(float64(result), float64(end-start)))

			// Remove all characters before the new
			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			// Slide the window since we have found a duplicate character
			start = duplicateIndex + 1
		}
		// Add the current character to hashmap
		charIndexMap[s[end]] = end
	}
	// Update the result for last window
	// For a input like "abc" the execution will not enter the above if statement for checking duplicates
	result = int(math.Max(float64(result), float64(n-start)))

	return result
}

func sliding_window_hashmap_02(s string) int {
	if len(s) == 0 {
        return 0
    }
    
    res := -1 << 63
    
    l := 0
    m := make(map[byte]int)
    m[s[0]]++
    
    for i := 1; i < len(s); i++ {
        m[s[i]]++
        
        for m[s[i]] > 1 {
            m[s[l]]--
            
            if m[s[l]] <= 0 {
                delete(m, s[l])
            }
            
            l++
        }
        
        if len(m) > res {
            res = len(m)
        }
    }
    
    if len(m) > res {
        res = len(m)
    }
    
    return res
}

/**
	dynamic programming
	Time complexity: O(n)O(n)O(n)
	Space complexity: O(1)O(1)O(1) (or O(alphabetSize)O(alphabetSize)O(alphabetSize) )
*/
func dp_solve(s string) int {
	rv := 0
    last := [256]int{}
    for i := range last {
        last[i] = -1
    }
    size := 0

    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    for i, b := range []byte(s) {
        size = min(size + 1, i - last[b])
        last[b] = i
        rv = rv + size - min(rv, size) // max(rv, size)
    }

    return rv
}

func main() {
	/**
		rabin karp
		https://stackoverflow.com/questions/72157919/are-kmp-and-rabin-karp-sliding-window-algorithms-for-pattern-matc
	*/
	input := "abcabcbb"
	input2 := "dundunqmop"
	input3 := "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur? At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere"
	input4 := "c"
	input5 := " "
	input6 := "au"
	input7 := "dvdf"
	fmt.Println(sliding_window_hashmap_02(input))
	fmt.Println(sliding_window_hashmap_02(input2))
	fmt.Println(sliding_window_hashmap_02(input3))
	fmt.Println(sliding_window_hashmap_02(input4)) // 1
	fmt.Println(sliding_window_hashmap_02(input5)) // 1
	fmt.Println(sliding_window_hashmap_02(input6)) // 2
	fmt.Println(sliding_window_hashmap_02(input7))
}
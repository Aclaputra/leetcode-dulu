package main

import "fmt"

// complexity time : o(n^2), space : o(1)
func brute_force_solve(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// complexity time : o(n), space : o(n)
func hashmap_solve(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target - currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}

func main() {
	input_num_01 := []int{7,2,13,11,8}
	input_target_01 := 24
	fmt.Println(brute_force_solve(input_num_01, input_target_01))
	fmt.Println(hashmap_solve(input_num_01, input_target_01))
}
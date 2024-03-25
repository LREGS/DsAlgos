/*128. Longest Consecutive Sequence
Medium
Topics
Companies
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.*/

package main

import (
	"fmt"
)

//we want to sort the list, and iterate over the list and subtract the ith + 1 index from the current number, and if the result == 1 append to consec list, else, consec list is complete

// need to use i instead and then do range (len(nums) - )
func longestConsecutive(nums []int) int {

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	//return max of result
	longest := 0

	for n := range set {

		if _, ok := set[n-1]; !ok {
			//if its a starting node
			seq := 1
			//this isnt it because it doesnt break - and we dont want it to loop while true because we are expecting it to be true - we need to now keep adding 1 to test start until it doesnt return true
			for {
				if _, ok := set[n+seq]; ok {
					seq++
					continue
				}
				longest = max(seq, longest)
				break

			}

		}

	}
	return longest
}

func main() {
	nums := []int{0, 1, 4, 3, 2, 5, 100, 102}
	answer := longestConsecutive(nums)

	fmt.Println(answer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func longestConsecutive(nums []int) int {

// 	set := make(map[int]bool)
// 	for _, num := range nums {
// 		set[num] = true
// 	}

// 	//return max of result
// 	longest := 0

// 	for i := 0; i < len(nums); i++ {
// 		testStart := nums[i] - 1

// 		if _, ok := set[testStart]; !ok {
// 			//if its a starting node
// 			seq := make([]int, 0)
// 			if _, ok := set[testStart]; ok {
// 				seq = append(seq, testStart)
// 			}
// 			//this isnt it because it doesnt break - and we dont want it to loop while true because we are expecting it to be true - we need to now keep adding 1 to test start until it doesnt return true
// 			for {
// 				testStart = testStart + 1
// 				_, ok := set[testStart]
// 				if ok {
// 					seq = append(seq, testStart)
// 				} else {
// 					break
// 				}
// 			}
// 			if longest < len(seq) {
// 				longest = len(seq)
// 			}

// 		}

// 	}
// 	return longest
// }

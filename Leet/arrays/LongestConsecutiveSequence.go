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

	consec := make([]int, 0)

	for i := 0; i < (len(nums) - 1); i++ {
		if nums[i+1]-nums[i] == 1 {
			consec = append(consec, nums[i])

		}
		if nums[i+1]-nums[i] != 1 {
			continue
		}
	}
	return len(consec)
}
func main() {
	nums := []int{5, 4, 3, 2, 1}
	answer := longestConsecutive(nums)

	fmt.Println(answer)
}

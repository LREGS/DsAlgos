/*128. Longest Consecutive Sequence
Medium
Topics
Companies
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.*/

package main

import (
	"fmt"
	"sort"
)

//we want to sort the list, and iterate over the list and subtract the ith + 1 index from the current number, and if the result == 1 append to consec list, else, consec list is complete

func longestConsecutive(nums []int) int {
	sort.Ints(nums)

	consec := make([]int, 0)

	for i, num := range nums {
		if i+1 > len(nums) {
			break
		}
		if nums[i+1]-num == 1 {
			consec = append(consec, num)
		}
		if nums[i+1]-num != 1 {
			break
		}
	}
	return len(consec)
}

func main() {
	nums := []int{4, 3, 2, 1}
	answer := longestConsecutive(nums)

	fmt.Println(answer)
}

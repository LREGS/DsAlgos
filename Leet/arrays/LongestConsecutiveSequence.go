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
	
	set := make(mape[int]bool)
	for num := range nums{
	    set[num] =true
	}



	consec := make(map[int][]int)



	for i := 0; i < (len(nums) - 1); i++ {
		test := nums[i] - 1
		_, ok := set[test]{
		swith ok{
		case false:
			consec[test] = append(consec[test], test)
	}
	return len(consec)
}
func main() {
	nums := []int{5, 4, 3, 2, 1}
	answer := longestConsecutive(nums)

	fmt.Println(answer)
}

package main

import (
	"fmt"
)

//  """Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

//  Example 1:

//  Input: nums = [1,1,1,2,2,3], k = 2
//  Output: [1,2]
//  Example 2:

//  Input: nums = [1], k = 1
//  Output: [1]"""

//same as last problem - add all the elements into a dictionary, return all elements from the dictionary into a list, sort the list, return top 2 elements

func topKFrequent(nums []int) {

	tokenMap := make(map[int][]int)

	for _, num := range nums {
		tokenMap[num] = append(tokenMap[num], num)
	}

	for _, value := range tokenMap {
		fmt.Println(len(value))
	}

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 8, 8, 8}
	topKFrequent(nums)

}

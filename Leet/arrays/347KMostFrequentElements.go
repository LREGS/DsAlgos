package main

import (
	"fmt"
	"sort"
)

//  """Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

//  Example 1:

//  Input: nums = [1,1,1,2,2,3], k = 2
//  Output: [1,2]
//  Example 2:

//  Input: nums = [1], k = 1
//  Output: [1]"""

//same as last problem - add all the elements into a dictionary, return all elements from the dictionary into a list, sort the list, return top 2 elements

func topKFrequent(nums []int, k int) []int {

	tokenMap := make(map[int][]int)

	for _, num := range nums {
		tokenMap[num] = append(tokenMap[num], num)
	}

	answers := make([][]int, 0)

	for _, value := range tokenMap {
		answers = append(answers, value)
	}

	lessFunc := func(i, j int) bool {
		return len(answers[j]) < len(answers[i])
	}

	sort.Slice(answers, lessFunc)

	selected := answers[:k]

	actualAnswer := make([]int, 0)

	for _, value := range selected {
		if len(actualAnswer) == k {
			break
		}
		actualAnswer = append(actualAnswer, value[1])

	}

	return actualAnswer

}

func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 5, 6, 6, 7, 7, 7, 8, 8, 8}
	answer := topKFrequent(nums, 3)
	fmt.Println(answer)

}

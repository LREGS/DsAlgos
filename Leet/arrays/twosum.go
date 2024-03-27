package main

import "fmt"

func twoSum(numbers []int, target int) []int {

	res := make([]int, 0)

	start := 0
	end := len(numbers) - 1

	for start < end {

		t := (numbers[start] + numbers[end])

		switch {
		case t == target:
			res = append(res, start+1, end+1)
			return res
		case t < target:
			start++
		case t > target:
			end = end - 1
		}

	}
	return res

}

// for i, num := range numbers {
// 	t := target - num
// 	for j := i + 1; j < len(numbers); j++{
// 		if numbers[j] == t {
// 			res = append(res, i+1, j+1)
// 			return res
// 		}
// 	}
// }
// return res

func main() {
	numbers := []int{1, 2}
	targer := 3
	res := twoSum(numbers, targer)
	fmt.Print(res)
}

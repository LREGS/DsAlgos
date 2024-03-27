package main

import "fmt"

func twoSum(numbers []int, target int) []int {

	res := make([]int, 0)

	for i, num := range numbers {
		t := target - num
		for j, num := range numbers[i:] {
			if num == t {
				// if i == j {
				// 	res = append(res, i, j+1)
				// 	return res
				// }
				res = append(res, i+1, j+1)
				return res
			}
		}
	}
	return res

}

func main() {
	numbers := []int{1, 2}
	targer := 3
	res := twoSum(numbers, targer)
	fmt.Print(res)
}

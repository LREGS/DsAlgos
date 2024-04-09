package main

import "fmt"

func threeSum(num []int) {

	// target := 0

	for i := 0; i < len(num); i++ {
		// answer := make([][]int, 0)

		y := i + 1
		z := (len(num) - 1)
		for y < z {
			// subAns := make([]int, 0)
			fmt.Print(z)
			// fmt.Printf("%d, %d, %d", num[i], num[y], num[z])
			y++
			z--

		}

	}

}

func main() {
	nums := []int{-1, 1, 0, 2, -3, -5, 6, 3}
	threeSum(nums)
}

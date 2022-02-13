package main

import "fmt"

func main() {
	fmt.Println(getTotalX([]int{2, 4}, []int{16, 32, 96}))
}

func getTotalX(a []int, b []int) int {
	// Write your code here

	var c int
	for i := 1; i <= 100; i++ {
		factor := true

		//check first
		for j := 0; j < len(a); j++ {
			if i%a[j] != 0 {
				factor = false
				break
			}
		}
		// check second
		for j := 0; j < len(b); j++ {
			if b[j]%i != 0 {
				factor = false
				break
			}
		}
		//update counter
		if factor {
			c++
			fmt.Println(i)
		}

	}
	return c
}

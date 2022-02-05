package main

import "fmt"

func main() {
	var arrSum []int
	arr := []int{1, 3, 5, 4, 2}
	//fmt.Println(arr[4-1])
	for i := 0; i < len(arr); i++ {
		a := arr[i] * arr[i+1]
		arrSum = append(arrSum, a)
		fmt.Println(a)
	}

	var sum = 0
	for _, i := range arrSum {
		sum += i
	}
	fmt.Println(sum)
}

package main

import "fmt"

func main() {
	countApplesAndOranges(7, 10, 4, 12, []int32{2, 3, -4}, []int32{3, -2, -4})
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var newApple []int32
	var newOrange []int32
	var appleCount int32
	var orangeCount int32

	for _, v := range apples {
		newApple = append(newApple, v+a)
	}
	for _, v := range oranges {
		newOrange = append(newOrange, v+b)
	}
	for _, v := range newApple {
		if v >= s && v <= t {
			appleCount++
		}
	}
	for _, v := range newOrange {
		if v >= s && v <= t {
			orangeCount++
		}
	}
	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}

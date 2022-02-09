package main

import "fmt"

func main() {
	fmt.Println(birthdayCakeCandles([]int32{82, 49, 82, 82, 41, 82, 15, 63, 38, 25}))
}

func birthdayCakeCandles(candles []int32) int32 {
	var count int32
	m := map[int32]int32{}
	for _, v := range candles {
		m[v]++
	}
	for _, v := range m {
		if v > count {
			count = v
		}
	}
	return count
}

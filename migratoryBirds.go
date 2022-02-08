package main

import "fmt"

func main() {
	fmt.Println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}))
}

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	m := map[int32]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v]++
		}
	}
	var key int32
	var val int
	for k, v := range m {
		if v > val {
			val = v
			key = k
		}
	}
	return key
}

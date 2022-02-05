package main

import "fmt"

func main() {
	RemoveDuplicates([]int{1, 2, 2, 23, 5, 6, 8, 9, 23})
}

func RemoveDuplicates(n []int) {
	arr := []int{}
	m := map[int]int{}
	for _, v := range n {
		if _, ok := m[v]; !ok {
			m[v] = v
			arr = append(arr, v)
		}
	}
	fmt.Println(arr)
}

package main

import "fmt"

func main() {
	fmt.Println(circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	result := []int32{}
	newres := []int32{}
	for i := int32(0); i < k; i++ {
		last := a[len(a)-1]
		a = a[:len(a)-1]
		result = append(result, last)
		result = append(result, a...)
		a = result
		if i != k-1 {
			result = nil
		}
	}
	for i := 0; i < len(queries); i++ {
		newres = append(newres, result[queries[i]])
	}
	return newres
}

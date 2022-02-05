package main

import (
	"fmt"
	"sort"
)

func main()  {
	k:= 4
	h:= []int{1,2,3,2}
	fmt.Println(hurdleRace(k,h))
}

func hurdleRace(k int, height []int) int {
	// Write your code here
	var n int
	sort.Ints(height)
	for i:=0; i<len(height);i++{
		var s int= height[len(height)-1]
		if k>s{
			return 0
		}
		if k<s{
			n= s-k
		}
	}
	return n


}
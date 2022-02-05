package main

import "fmt"

func main()  {
	ar:= []int32{1,2,3,4,5,6}
	k:= int32(5)
	n:= 6 // len of ar[]
	fmt.Println(divisibleSumPairs(n,k,ar))
}

func divisibleSumPairs(n int, k int32, ar []int32) int32 {
	// Write your code here
	var count int32= 0
	for i:=0; i<n; i++{
		for j:=i+1; j<n; j++{
			if (ar[i]+ar[j])%k==0{
				count++
			}
		}
	}
	return count

}
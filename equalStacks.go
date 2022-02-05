package main

import "fmt"

func main() {
	h1:= []int{1,1,1,2,3}
	var sumed1 []int
	var sumed2 []int
	var sumed3 []int
	var rev1 []int
	var rev2 []int
	var rev3 []int
	//var sum []int
	h2:= []int{2,3,4}
	h3:= []int{1,4,1,1}
	s1:=0
	s2:=0
	s3:=0
	for _,i:= range h1{
		s1 += i
		sumed1= append(sumed1,s1)

	}
	for _,i:= range h2{
		s2 += i
		sumed2= append(sumed2,s2)

	}
	for _,i:= range h3{
		s3 += i
		sumed3= append(sumed3,s3)

	}
	sum1:=0
	sum2:=0
	sum3:=0
	for i,j:= range sumed1{
		rev1 = append(rev1, sumed1[len(sumed1)-1-i])
		sum1 += j
	}
	for i,j:= range sumed2{
		rev2 = append(rev2, sumed2[len(sumed2)-1-i])
		sum2 += j
	}
	for i,j:= range sumed3{
		rev3 = append(rev3, sumed3[len(sumed3)-1-i])
		sum3 += j
	}
	smallest:= min(sum1, sum2,sum3)
	//if len(rev1)


	fmt.Println(smallest)
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(sum3)
	//fmt.Println(sumed2)
	//fmt.Println(sum3)
	//fmt.Println(sumed3)
}
//func min(a ,b, c int) int {
//	if a <= b && a <= c {
//		return a
//	} else if b <= a && b <= c {
//		return b
//	}
//	return c
//}
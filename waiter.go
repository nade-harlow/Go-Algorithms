package main

import "fmt"

func main() {
	nu:= []int32{2,3,4,5,6,7}
	q:= 2
	fmt.Println(waiter(nu,q))
}


func waiter(number []int32, q int) []int32 {
	// Write your code here
	var A1 []int32
	var B1 []int32
	var answer []int32
	var prime =[]int32{2,3,5,7,11,13}
	var reversed []int32
	for i:= range number{
		reversed = append(reversed, number[len(number)-1-i])
	}
	//for i:=0; i<q; i++{
		for _,j:= range reversed{
			for _,v:= range prime{
				if j % v ==0{
					B1= append(B1, j)
				}else{
					A1= append(A1, j)
				}
			}
		}
		//for i:= range B1{
		//	answer = append(answer, B1[len(B1)-1-i])
		//}
		//if len(A1)== 0{
		//	return answer
		//}
	//}
	fmt.Println(B1)
	fmt.Println(A1)

	return answer
}



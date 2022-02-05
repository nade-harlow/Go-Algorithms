package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a:= []string{ "1 2","2","1 20","3"}
	fmt.Println(getMax(a))


}

func getMax(operations []string) []int32 {
	// Write your code here
	stack:= []int32{}
	maxi:= []int32{}
	for _,j:=range operations{
		if len(j)>1{
			spite:= strings.Split(j," ")
			num,_:= strconv.Atoi(spite[1])
			stack= append(stack, int32(num))
		}else{
			if j== "2"{
				stack= stack[:len(stack)-1]
			}

			if j == "3"{
				var max int32= 0
				for _,j:= range stack{
					if max<j{
						max= j
					}
				}
				maxi= append(maxi, max)
			}
		}
	}
	return maxi
}
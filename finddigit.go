package main


import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println(findDigits(1012))
}
func findDigits(n int) int {
	// Write your code here
	var count int
	var st []string
	var converted []int
	str:= strconv.Itoa(n)
	st = strings.Split(str, "")
	for _, i := range st {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		converted = append(converted, j)

	}
	fmt.Println(len(converted))
	fmt.Println(converted)
	fmt.Println(n)
	for i,_:=range converted{
		if converted[i]==0{
			continue
		}
		if n%converted[i]==0{
			count++
		}
	}
	return count
}

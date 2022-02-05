package main

import (
	"fmt"
	"strings"
)

func main()  {
	s1:= "hello"
	s2:= "ll"
	fmt.Println(twoStrings(s1,s2))
}
func twoStrings(s1 string, s2 string) string {
	// Write your code here
	if strings.ContainsAny(s1,s2){
		return "YES"
	}else{
		return "NO"
	}


}
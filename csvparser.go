//package main
//
//import "fmt"
//
//func main() {
//	fmt.Println(Numerals(1))
//}
//
//
//func Numerals(num int) string {
//	var roman string
//	if num == 1{
//		roman = "I"
//	}
//	if num == 2{
//		roman = "II"
//	}
//	if num == 3{
//		roman = "III"
//	}
//	if num == 4{
//		roman = "IV"
//	}
//	if num == 5{
//		roman = "V"
//	}
//	if num == 6{
//		roman = "VI"
//	}
//
//	return roman
//}
package main

import (
	"fmt"
	"strings"
)

//var GlobalFlag string
const (
	Write = iota
	Read
	Execute
)

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func main() {
	var r []string

	denominations:= []string{"BRA", "KSA", "USA", "BRA", "JPN", "PHL", "KSA", "UAE", "UAE", "JPN"}
		r= append(r, denominations[0])
		v:=removeDuplicateStr(denominations)
		//sort.Strings(denominations)

		//fmt.Println(r)
	//temp:= r[0]
		//for i:= 0; i<len(denominations);i++{
		//	for j:= 1; j<len(denominations)-1;j++{
		//		if temp == denominations[j]{
		//			t= append(t, denominations[j])
		//		}
		//
		//	}
		//	temp= denominations[1]
		//}
		//fmt.Println(r)
		//fmt.Println(denominations)
	//count:= 0
	//target:= 100
	//for _, v := range denominations {
	//	quot := target / v
	//	remain := target % v
	//	target = remain
	//	count += quot
	//}
	fmt.Println(strings.Join(v," "))
}

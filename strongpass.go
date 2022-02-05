package main

import (
	"fmt"
	"strings"
)

func main()  {

	count:= 0
	password:= "ashq"
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"

	pw:=strings.Split(password, "")
	//n:=strings.Split(numbers, "")
	//lc:=strings.Split(lower_case, "")
	//uc:=strings.Split(upper_case, "")
	//sc:=strings.Split(special_characters, "")
for i:=range pw{
	if strings.ContainsAny(password, numbers){
		count +=i
	}
	if strings.ContainsAny(password, lower_case){
		count +=i
	}
	if strings.ContainsAny(password,upper_case){
		count +=i
	}
	if strings.ContainsAny(password, special_characters){
		count +=i
	}



}

	fmt.Println(count)
	//fmt.Print(sc)

}


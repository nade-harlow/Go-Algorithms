package main

import "fmt"

func main() {
	ReverseString("james")
}

func ReverseString(s string) {
	newStr := ""
	for i := 0; i < len(s); i++ {
		newStr += string(s[len(s)-1-i])
	}
	fmt.Println(newStr)
}

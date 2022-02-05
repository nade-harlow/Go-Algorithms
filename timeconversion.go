package main

import "fmt"

func main(){
	var h, m, s int
	var l string
	fmt.Scanf("%d:%d:%d%s", &h,&m,&s,&l)

	if l== "AM" || l== "am" && h == 12{
		h=0
	}
	if l== "PM" || l== "pm" && h != 12{
		h +=12
	}
	fmt.Printf("%02d:%02d:%02d",h,m,s)
}

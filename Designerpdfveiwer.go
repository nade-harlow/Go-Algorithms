package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	var g []int
	h:= []int{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	a:=[]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	word:= "zaba"
	w:= strings.Split(word, "")

	for i,v:=range a{
		for j:= range w{
			if w[j]==v{
				for r,d:=range h{
					if i==r{
						g= append(g,d)
					}
				}
			}
		}
	}
	sort.Ints(g)
	ff:= g[len(g)-1]

	fmt.Println(ff*len(g))

}

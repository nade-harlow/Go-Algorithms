package main
import "fmt"

func main()  {
	a:= []int32{-1, -3, 4, 2}
	fmt.Println(angryProfessor(3,a))
}

func angryProfessor(k int32, a []int32) string {
	var count int32
	for _,v:=range a{
		if v<=0{
			count++

		}
	}
	if k>count{
		return "YES"
	}else{
		return "NO"
	}
}

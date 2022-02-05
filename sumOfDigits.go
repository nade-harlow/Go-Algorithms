package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	var nn string
	var nnn []string
	var nnm []int64
	n:= int64(34)
	nn= strconv.FormatInt(n, 10)
	//nn=string(n)
	nnn= strings.Split(nn,"")
	fmt.Println(nnn)//string
	//fmt.Printf("%T",nnn)
	//fmt.Print(n)
	for _,v:=range nnn{
		if num,err:= strconv.ParseInt(v, 10, 8); err==nil{
			nnm = append(nnm, num)


		}

	}
	fmt.Println(nnm)//int
	//fmt.Printf("%T",nnm)
	var sum int64
	var summ []int64
	for _,v:=range nnm{
		sum += v
	}
	summ= append(summ,sum)
	//fmt.Println(summ)
	var delim string
	var del string
	for i:=range summ{
		if summ[i]>0{
			del=strings.Trim(strings.Replace(fmt.Sprint(summ), " ", delim, -1), "[]")
			fmt.Println(del)
			fmt.Printf("%T", del)
		}
	}



}

package main

import (
	"fmt"
	"math/big"
)
func main(){
	var n int64
	fmt.Scan(&n)
	//for i:=0; i<n; n--{
	//	n*= n-1
	//}
	factorial(n)
}
// 4*3*2*1

func factorial(n int64) {
	x := big.NewInt(int64(1))

	for i:= int64(1); i<=n; i++ {
		x = x.Mul(x, big.NewInt(i))
	}
	fmt.Println(x)
}
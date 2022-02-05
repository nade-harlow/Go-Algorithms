package main

import "fmt"

func main() {
	FizzBuzz(15)
}

func FizzBuzz(n int) {
	//var arr []string
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if (i+1)%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if (i+1)%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i + 1)
	}

}

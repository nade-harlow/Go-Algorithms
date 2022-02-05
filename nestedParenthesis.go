package main

func main() {
	//fmt.Println(Nested("(()())"))
}

func Nested(s string) bool {
	myMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []rune
	for _, v := range s {
		if _, ok := myMap[v]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if myMap[top] != v {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

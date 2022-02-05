package main

import "fmt"

func main() {
	h1:= []int{3, 2, 1, 1, 1}
	h2:= []int{4, 3, 2}
	h3:= []int{1, 1, 4, 1}
	var rev1 []int
	var rev2 []int
	var rev3 []int

	sum1:=0
	sum2:=0
	sum3:=0
	for i,j:= range h1{
		rev1 = append(rev1, h1[len(h1)-1-i])
		sum1 += j
	}
	for i,j:= range h2{
		rev2 = append(rev2, h2[len(h2)-1-i])
		sum2 += j
	}
	for i,j:= range h3{
		rev3 = append(rev3, h3[len(h3)-1-i])
		sum3 += j
	}
	fmt.Println(rev1)
	fmt.Println(rev2)
	fmt.Println(rev3)
	//fmt.Println(sum1)
	//fmt.Println(sum2)
	//fmt.Println(sum3)
	m:= min(sum1,sum2,sum3)
		for sum1 >= m{
			sum1 -= rev1[len(rev1)-1]
			rev1= rev1[:len(rev1)-1]

		}
		for sum2 >= m {
			sum2 -= rev2[len(rev2)-1]
			rev2 = rev2[:len(rev2)-1]
		}
		for sum3 >= m {
			sum3 -= rev3[len(rev3)-1]
			rev3 = rev3[:len(rev3)-1]
		}

	if sum1==sum2 && sum2==sum3{
		fmt.Println(sum1)
	}else {
		fmt.Println(0)
	}
	//fmt.Println(m)
	//fmt.Println(sum1)
	//fmt.Println(sum2)
	//fmt.Println(sum3)
	//fmt.Println(rev1)
	//fmt.Println(rev2)
	//fmt.Println(rev3)
}

func min(a ,b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

/*
def max_length(h1, h2, h3):
    l1, l2, l3 = 0, 0, 0
    for each in h1:
        l1 = l1 + each
    for each in h2:
        l2 = l2 + each
    for each in h3:
        l3 = l3 + each
    while l1 !=0 and l2 != 0 and l3 != 0 and (l1!=l2 or l2!=l3):
        if max(l1, l2, l3) == l1:
            l1 = l1 - h1[0]
            h1.pop(0)
        elif max(l1, l2, l3) == l2:
            l2 = l2 - h2[0]
            h2.pop(0)
        else:
            l3 = l3 - h3[0]
            h3.pop(0)
    else:
        if l1==l2 and l2==l3:
            return l1
        else:
            return 0


def equalStacks(h1, h2, h3):
    # Write your code here
    s1, s2, s3 = map(sum, (h1, h2, h3))
    while h1 and h2 and h3:
        m = min(s1, s2, s3)
        while s1 > m: s1 -= h1.pop(0)
        while s2 > m: s2 -= h2.pop(0)
        while s3 > m: s3 -= h3.pop(0)
        if s1 == s2 == s3: return s1
    return 0
 */
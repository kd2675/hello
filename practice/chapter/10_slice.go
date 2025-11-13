package chapter

import "fmt"

func Slice() {
	var a = []int{1, 2, 3}
	fmt.Println(a)
	a[0] = 10
	fmt.Println(a)
	a = append(a, 4)
	fmt.Println(a)
	a = append(a, 5, 6)
	fmt.Println(a)
	a = append(a, []int{7, 8, 9}...)
	fmt.Println(a)
	a = append(a, []int{10, 11, 12}...)

	b := make([]int, 10)
	fmt.Println(b)

	var c = append(a, b...)
	fmt.Println(c)

	var d = a[1:]
	fmt.Println(d)
}

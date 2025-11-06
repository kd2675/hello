package practice

import "fmt"

func Variable() {
	var a int
	a = 10
	fmt.Println(a)

	var b, c int
	b, c = 1, 2
	fmt.Println(b, c)

	var d, e, f = 1, 2, 3
	fmt.Println(d, e, f)

	var g = "hello_g"
	fmt.Println(g)

	const h = "hello_h"
	fmt.Println(h)

	i := "hello_i"
	fmt.Println(i)
}

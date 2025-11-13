package chapter

import "fmt"

func DataType() {
	a := `simple_a\n`
	fmt.Println(a)
	b := "simple_b\n"
	fmt.Println(b)
	c := 10
	fmt.Println(c)
	d := 10.1
	fmt.Println(d)
	e := true
	fmt.Println(e)
	f := false
	fmt.Println(f)

	var g int = 100
	var h = uint(g)
	var i = float32(g)
	fmt.Println(g, h, i)

	j := []byte("hello_j")
	var n = []byte("hello_n")
	k := string(j)
	l := []string{"hello_l", "hello_l2"}
	m := map[string]string{"hello_m": "hello_m2"}
	m["test"] = "test2"
	fmt.Println(j, k, l, m)
	fmt.Println(n)
}

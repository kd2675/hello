package chapter

import "fmt"

type user struct {
	Name string
	Age  int
}

func newUser() user {
	u := user{
		Name: "b",
		Age:  20,
	}

	return u
}

func newUserPointer() *user {
	u := user{
		Name: "c",
		Age:  30,
	}

	return &u
}

func (u user) Say() {
	fmt.Println("hello")
}

func Struct() {
	u := user{
		Name: "a",
		Age:  1,
	}

	fmt.Println(u)

	u.Age = 2
	fmt.Println(u)

	u.Say()

	f := newUser()
	fmt.Println(f)

	f.Age = 10
	fmt.Println(f)

	g := newUserPointer()
	fmt.Println(*g)

	fmt.Println(*g)

	dic := newDict()
	dic.data[1] = "a"
	fmt.Println(dic.data)

	dic2 := newDict2()
	dic2.data[2] = "b"
	fmt.Println(dic2.data)

	dic.add(3, "c")
	dic.sout(3)
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func (d *dict) add(key int, value string) {
	d.data[key] = value
}

func newDict2() dict {
	d := dict{}
	d.data = map[int]string{}
	return d
}

func (d dict) sout(key int) {
	fmt.Println(d.data[key])
}

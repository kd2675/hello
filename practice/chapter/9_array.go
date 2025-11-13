package chapter

func Array() {
	var a [10]int
	a[0] = 10
	println(a[0])

	var b = [3]int{1, 2, 3}
	println(b[0])

	var c = [...]int{1, 2, 3}
	println(c[0])
}

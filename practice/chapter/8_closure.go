package chapter

func Closure() {
	plus := plusVal()

	plus()
	plus()
	plus()
}

func plusVal() func() {
	i := 0

	return func() {
		i++
		println(i)
	}
}

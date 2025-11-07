package practice

func AnonFunction() {
	printer := func(msgs ...string) {
		msg := ""
		for _, v := range msgs {
			msg += v
		}
		println(msg)
	}

	printer("hello", "world")
	printer("hello", "world", "hello")
	printer()

	printer2 := func(i string, j string) string {
		return i + j
	}
	i := add(printer2, "test", "Play")
	println(i)
}

func add(f func(string, string) string, a, b string) string {
	result := f(a, b)
	return result
}

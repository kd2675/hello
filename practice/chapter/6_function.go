package chapter

func Function() {
	msg := "hi"
	say(&msg)
	println(msg)

	says("hello", "world")
	says("hello", "world", "hello")

	println(sayReturn())

	count, total := sum(1, 7, 3, 5, 9)
	println(count, total)
}

func say(msg *string) {
	println(*msg)
	*msg = "hello"
}

func says(msg ...string) {
	for _, v := range msg {
		println(v)
	}
}

func sayReturn() string {
	return "hello"
}

func sum(nums ...int) (int, int) {
	s := 0     // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

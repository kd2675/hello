package chapter

var practiceMap map[string]string

func init() {
	practiceMap = make(map[string]string)
}

func Package() {
	practiceMap["a"] = "a"
	println(practiceMap["a"])
}

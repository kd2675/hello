package practice

func Conditional() {
	if i := 0; i < 1 {
		println("if : ", i)
	}

	switch i := 0; i {
	case 0:
		println("switch: ", i)
	default:
		println("default: ", i)
	}

	gradeInt(90)
	gradeInt(80)
	gradeInt(70)
	gradeInterface(90)
	gradeInterface("90")
}

func gradeInt(score int) {
	switch {
	case score >= 90:
		println("A")
		fallthrough
	case score >= 80:
		println("B")
	default:
		println("C")
	}
}

func gradeInterface(score interface{}) {
	switch score.(type) {
	case int:
		println("int")
	case string:
		println("string")
	default:
		println("unknown")
	}
}

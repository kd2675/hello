package practice

func Loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}
	println(sum)

	for sum < 20 {
		sum++
	}
	println(sum)

	for {
		break
	}
	println("end")

	names := []string{"a", "b", "c"}
	for _, name := range names {
		println(name)
	}

	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)
END:
	println("end")

L1:
	for {
		if sum == 20 {
			break L1
		}
	}

	println("OK")
}

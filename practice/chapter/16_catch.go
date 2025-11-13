package chapter

import "fmt"

// recover
// 함수 f의 depth가 n이고, f에서 panic이 발생하여 복구하고 싶다면 recover는 지연 호출되어야 하고 depth는 n+1이 되어야 한다.
func Catch() {
	defer b()
	defer c()

	panic("panic!")

}

func c() {
	r := recover()
	fmt.Println("recoverd: ", r)
	fmt.Println("a called")
}

func b() {
	fmt.Println("b called")
}

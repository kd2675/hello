package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine() (string, error) {
	fmt.Println("덧셈할 문자열을 입력해 주세요.")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	return input, nil
}

func PrintResult(result int64) {
	fmt.Printf("결과 : %d\n", result)
}

package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func ReadCarNames() (string, error) {
	fmt.Println("경주할 자동차 이름을 입력하세요.(이름은 쉼표(,) 기준으로 구분)")

	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return scanner.Text(), nil
}

func ReadRounds() (string, error) {
	fmt.Println("시도할 횟수는 몇 회인가요?")

	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return scanner.Text(), nil
}

func ParseCarNames(input string) []string {
	names := strings.Split(input, ",")
	result := make([]string, 0, len(names))

	for _, name := range names {
		trimmed := strings.TrimSpace(name)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

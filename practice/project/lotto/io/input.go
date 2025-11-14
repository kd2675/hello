package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func InputPurchaseAmount() (string, error) {
	fmt.Println("구입금액을 입력해 주세요.")

	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return scanner.Text(), nil
}

func InputWinningNumber() (string, error) {
	fmt.Println()
	fmt.Println("당첨 번호를 입력해 주세요.")

	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return scanner.Text(), nil
}

func InputBonusNumber() (string, error) {
	fmt.Println()
	fmt.Println("보너스 번호를 입력해 주세요.")

	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return scanner.Text(), nil
}

func ParseNumbers(input string) ([]int, error) {
	parts := strings.Split(input, ",")
	numbers := make([]int, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}

		num, err := strconv.Atoi(trimmed)
		if err != nil {
			return nil, fmt.Errorf("숫자 형식이 올바르지 않습니다: %s", trimmed)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

package model

import (
	"errors"
	"strconv"
	"strings"
)

type RoundCount struct {
	value int
}

func NewRoundCount(input string) (*RoundCount, error) {
	if input == "" {
		return nil, errors.New("시도 횟수를 입력해주세요")
	}

	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return nil, errors.New("시도 횟수를 입력해주세요")
	}

	value, err := strconv.Atoi(trimmed)
	if err != nil {
		return nil, errors.New("시도 횟수는 숫자여야 합니다")
	}

	if value <= 0 {
		return nil, errors.New("시도 횟수는 양수여야 합니다")
	}

	return &RoundCount{
		value: value,
	}, nil
}

func (r *RoundCount) Value() int {
	return r.value
}

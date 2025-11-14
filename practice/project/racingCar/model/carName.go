package model

import (
	"errors"
	"strings"
	"unicode/utf8"
)

const maxNameLength = 5

type CarName struct {
	value string
}

func NewCarName(value string) (*CarName, error) {
	if err := validateCarName(value); err != nil {
		return nil, err
	}

	return &CarName{
		value: strings.TrimSpace(value),
	}, nil
}

func (c *CarName) Value() string {
	return c.value
}

func validateCarName(value string) error {
	if value == "" {
		return errors.New("자동차 이름은 비어있을 수 없습니다")
	}

	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return errors.New("자동차 이름은 공백만으로 구성될 수 없습니다")
	}

	if utf8.RuneCountInString(trimmed) > maxNameLength {
		return errors.New("자동차 이름은 5자 이하여야 합니다")
	}

	return nil
}

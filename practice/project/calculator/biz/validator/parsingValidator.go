package validator

import (
	"errors"
)

type ParsingValidator interface {
	Validate(numbers []int64) error
}

type parsingValidatorImpl struct{}

func NewParsingValidator() ParsingValidator {
	return &parsingValidatorImpl{}
}

func (v *parsingValidatorImpl) Validate(numbers []int64) error {
	for _, num := range numbers {
		if num < 0 {
			return errors.New("음수는 허용되지 않습니다")
		}
	}
	return nil
}

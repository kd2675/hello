package model

import (
	"errors"
	"sort"
)

const (
	lottoSize   = 6
	minLottoNum = 1
	maxLottoNum = 45
)

type Lotto struct {
	numbers []int
}

func NewLotto(numbers []int) (*Lotto, error) {
	if err := validate(numbers); err != nil {
		return nil, err
	}
	if err := validateRange(numbers); err != nil {
		return nil, err
	}
	if err := validateDuplicate(numbers); err != nil {
		return nil, err
	}

	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)

	return &Lotto{numbers: sorted}, nil
}

func (l *Lotto) Numbers() []int {
	result := make([]int, len(l.numbers))
	copy(result, l.numbers)
	return result
}

func (l *Lotto) Contains(number int) bool {
	for _, n := range l.numbers {
		if n == number {
			return true
		}
	}
	return false
}

func validate(numbers []int) error {
	if len(numbers) != lottoSize {
		return errors.New("[ERROR] 로또 번호는 6개여야 합니다")
	}
	return nil
}

func validateRange(numbers []int) error {
	for _, number := range numbers {
		if number < minLottoNum || number > maxLottoNum {
			return errors.New("[ERROR] 로또 번호는 1-45 사이 숫자이어야합니다")
		}
	}
	return nil
}

func validateDuplicate(numbers []int) error {
	seen := make(map[int]bool)
	for _, number := range numbers {
		if seen[number] {
			return errors.New("[ERROR] 로또 번호는 중복될 수 없습니다")
		}
		seen[number] = true
	}
	return nil
}

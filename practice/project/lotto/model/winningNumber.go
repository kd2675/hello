package model

import "errors"

type WinningNumber struct {
	numbers []int
}

func NewWinningNumber(numbers []int) (*WinningNumber, error) {
	if len(numbers) == 0 {
		return nil, errors.New("당첨 번호를 입력해주세요")
	}

	if err := validate(numbers); err != nil {
		return nil, err
	}
	if err := validateRange(numbers); err != nil {
		return nil, err
	}
	if err := validateDuplicate(numbers); err != nil {
		return nil, err
	}

	return &WinningNumber{numbers: numbers}, nil
}

func (w *WinningNumber) CountMatch(lotto *Lotto) int {
	count := 0
	lottoNumbers := lotto.Numbers()

	for _, lottoNum := range lottoNumbers {
		if w.contains(lottoNum) {
			count++
		}
	}

	return count
}

func (w *WinningNumber) Contains(number int) bool {
	return w.contains(number)
}

func (w *WinningNumber) contains(number int) bool {
	for _, n := range w.numbers {
		if n == number {
			return true
		}
	}
	return false
}

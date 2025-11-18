package model

import "errors"

type BonusNumber struct {
	number int
}

func NewBonusNumber(number int, winningNumber *WinningNumber) (*BonusNumber, error) {
	if err := validateBonusRange(number); err != nil {
		return nil, err
	}
	if err := validateBonusDuplicate(number, winningNumber); err != nil {
		return nil, err
	}

	return &BonusNumber{number: number}, nil
}

func (b *BonusNumber) Match(lotto *Lotto) bool {
	return lotto.Contains(b.number)
}

func validateBonusRange(number int) error {
	if number < minLottoNum || number > maxLottoNum {
		return errors.New("보너스 번호는 1-45 사이 숫자이어야합니다")
	}
	return nil
}

func validateBonusDuplicate(number int, winningNumber *WinningNumber) error {
	if winningNumber.Contains(number) {
		return errors.New("보너스 번호는 당첨 번호와 중복될 수 없습니다")
	}
	return nil
}

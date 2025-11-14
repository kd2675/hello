package service

import (
	"lotto/io"
	"lotto/model"
	"strconv"
)

type InputService interface {
	InputPurchaseAmount() (int, error)
	InputWinningNumber() (*model.WinningNumber, error)
	InputBonusNumber(winningNumber *model.WinningNumber) (*model.BonusNumber, error)
}

type inputServiceImpl struct{}

func NewInputService() InputService {
	return &inputServiceImpl{}
}

func (s *inputServiceImpl) InputPurchaseAmount() (int, error) {
	for {
		input, err := io.InputPurchaseAmount()
		if err != nil {
			return 0, err
		}

		purchaseAmount, err := model.NewPurchaseAmount(input)
		if err != nil {
			io.PrintErrorMessage(err)
			continue
		}

		return purchaseAmount.Price(), nil
	}
}

func (s *inputServiceImpl) InputWinningNumber() (*model.WinningNumber, error) {
	for {
		input, err := io.InputWinningNumber()
		if err != nil {
			return nil, err
		}

		numbers, err := io.ParseNumbers(input)
		if err != nil {
			io.PrintErrorMessage(err)
			continue
		}

		winningNumber, err := model.NewWinningNumber(numbers)
		if err != nil {
			io.PrintErrorMessage(err)
			continue
		}

		return winningNumber, nil
	}
}

func (s *inputServiceImpl) InputBonusNumber(winningNumber *model.WinningNumber) (*model.BonusNumber, error) {
	for {
		input, err := io.InputBonusNumber()
		if err != nil {
			return nil, err
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			io.PrintErrorMessage(err)
			continue
		}

		bonusNumber, err := model.NewBonusNumber(number, winningNumber)
		if err != nil {
			io.PrintErrorMessage(err)
			continue
		}

		return bonusNumber, nil
	}
}

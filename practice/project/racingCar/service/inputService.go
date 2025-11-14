package service

import (
	"racingCar/io"
	"racingCar/model"
)

type InputService interface {
	GetCarNames() ([]string, error)
	GetRoundCount() (*model.RoundCount, error)
}

type inputServiceImpl struct{}

func NewInputService() InputService {
	return &inputServiceImpl{}
}

func (s *inputServiceImpl) GetCarNames() ([]string, error) {
	input, err := io.ReadCarNames()
	if err != nil {
		return nil, err
	}

	return io.ParseCarNames(input), nil
}

func (s *inputServiceImpl) GetRoundCount() (*model.RoundCount, error) {
	input, err := io.ReadRounds()
	if err != nil {
		return nil, err
	}

	return model.NewRoundCount(input)
}

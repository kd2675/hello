package service

import "racingCar/utils"

type MoveService interface {
	GetMoveValues(count int) []int
}

type randomMoveService struct{}

func NewRandomMoveService() MoveService {
	return &randomMoveService{}
}

func (s *randomMoveService) GetMoveValues(count int) []int {
	result := make([]int, count)

	for i := 0; i < count; i++ {
		result[i] = utils.GenerateRandomValue()
	}

	return result
}

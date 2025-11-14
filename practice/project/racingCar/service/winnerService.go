package service

import "racingCar/model"

type WinnerService interface {
	DetermineWinners(carList *model.CarList) []string
}

type winnerServiceImpl struct{}

func NewWinnerService() WinnerService {
	return &winnerServiceImpl{}
}

func (s *winnerServiceImpl) DetermineWinners(carList *model.CarList) []string {
	maxPosition := s.findMaxPosition(carList)
	return s.findWinnerNames(carList, maxPosition)
}

func (s *winnerServiceImpl) findMaxPosition(carList *model.CarList) int {
	max := 0
	for _, car := range carList.Cars() {
		if car.Position() > max {
			max = car.Position()
		}
	}
	return max
}

func (s *winnerServiceImpl) findWinnerNames(carList *model.CarList, maxPosition int) []string {
	winners := make([]string, 0)
	for _, car := range carList.Cars() {
		if car.Position() == maxPosition {
			winners = append(winners, car.Name())
		}
	}
	return winners
}

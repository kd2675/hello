package service

import "racingCar/model"

type RaceService interface {
	ExecuteRound(carList *model.CarList, randomValues []int)
}

type raceServiceImpl struct{}

func NewRaceService() RaceService {
	return &raceServiceImpl{}
}

func (s *raceServiceImpl) ExecuteRound(carList *model.CarList, randomValues []int) {
	carList.MoveAll(randomValues)
}

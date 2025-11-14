package controller

import (
	"fmt"
	"racingCar/io"
	"racingCar/model"
	"racingCar/service"
)

type RacingGameController struct {
	inputService  service.InputService
	moveService   service.MoveService
	raceService   service.RaceService
	winnerService service.WinnerService
}

func NewRacingGameController(
	inputService service.InputService,
	moveService service.MoveService,
	raceService service.RaceService,
	winnerService service.WinnerService,
) *RacingGameController {
	return &RacingGameController{
		inputService:  inputService,
		moveService:   moveService,
		raceService:   raceService,
		winnerService: winnerService,
	}
}

func (c *RacingGameController) StartGame() error {
	carNames, err := c.inputService.GetCarNames()
	if err != nil {
		return fmt.Errorf("자동차 이름 입력 오류: %w", err)
	}

	carList, err := model.NewCarList(carNames)
	if err != nil {
		return fmt.Errorf("자동차 목록 생성 오류: %w", err)
	}

	roundCount, err := c.inputService.GetRoundCount()
	if err != nil {
		return fmt.Errorf("시도 횟수 입력 오류: %w", err)
	}

	if err := c.executeRounds(carList, roundCount.Value()); err != nil {
		return err
	}

	c.printWinners(carList)
	return nil
}

func (c *RacingGameController) executeRounds(carList *model.CarList, roundCount int) error {
	io.PrintResultMessage()

	for i := 0; i < roundCount; i++ {
		if err := c.executeRound(carList); err != nil {
			return err
		}
		io.PrintRoundResult(carList)
	}

	return nil
}

func (c *RacingGameController) executeRound(carList *model.CarList) error {
	carCount := len(carList.Cars())
	randomValues := c.moveService.GetMoveValues(carCount)
	c.raceService.ExecuteRound(carList, randomValues)
	return nil
}

func (c *RacingGameController) printWinners(carList *model.CarList) {
	winners := c.winnerService.DetermineWinners(carList)
	io.PrintWinners(winners)
}

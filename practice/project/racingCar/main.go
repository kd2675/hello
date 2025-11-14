package main

import (
	"fmt"
	"os"
	"racingCar/controller"
	"racingCar/service"
)

func main() {
	inputService := service.NewInputService()
	moveService := service.NewRandomMoveService()
	raceService := service.NewRaceService()
	winnerService := service.NewWinnerService()

	gameController := controller.NewRacingGameController(
		inputService,
		moveService,
		raceService,
		winnerService,
	)

	if err := gameController.StartGame(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package io

import (
	"fmt"
	"racingCar/model"
	"strings"
)

func PrintResultMessage() {
	fmt.Println()
	fmt.Println("실행 결과")
}

func PrintRoundResult(carList *model.CarList) {
	for _, car := range carList.Cars() {
		printCarStatus(car)
	}
	fmt.Println()
}

func printCarStatus(car *model.Car) {
	fmt.Printf("%s : %s\n", car.Name(), car.PositionString())
}

func PrintWinners(winners []string) {
	fmt.Printf("최종 우승자 : %s\n", strings.Join(winners, ", "))
}

package main

import (
	"fmt"
	"lotto/controller"
	"lotto/service"
	"os"
)

func main() {
	// 서비스 생성
	inputService := service.NewInputService()
	lottoService := service.NewLottoService()
	winningService := service.NewWinningService()

	// 컨트롤러 생성
	lottoController := controller.NewLottoController(
		inputService,
		lottoService,
		winningService,
	)

	// 게임 시작
	if err := lottoController.Start(); err != nil {
		fmt.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}

package controller

import (
	"lotto/io"
	"lotto/model"
	"lotto/service"
)

type LottoController struct {
	inputService   service.InputService
	lottoService   service.LottoService
	winningService service.WinningService
}

func NewLottoController(
	inputService service.InputService,
	lottoService service.LottoService,
	winningService service.WinningService,
) *LottoController {
	return &LottoController{
		inputService:   inputService,
		lottoService:   lottoService,
		winningService: winningService,
	}
}

func (c *LottoController) Start() error {
	// 구입 금액 입력
	purchaseAmount, err := c.inputService.InputPurchaseAmount()
	if err != nil {
		return err
	}

	// 지갑 생성
	wallet, err := model.NewWallet(purchaseAmount)
	if err != nil {
		io.PrintErrorMessage(err)
		return err
	}

	// 로또 구매
	if err := c.lottoService.BuyLottos(wallet); err != nil {
		io.PrintErrorMessage(err)
		return err
	}

	// 당첨 번호 입력
	winningNumber, err := c.inputService.InputWinningNumber()
	if err != nil {
		return err
	}

	// 보너스 번호 입력
	bonusNumber, err := c.inputService.InputBonusNumber(winningNumber)
	if err != nil {
		return err
	}

	// 당첨 확인 및 수익률 계산
	c.winningService.WinningLottos(wallet, winningNumber, bonusNumber)

	return nil
}

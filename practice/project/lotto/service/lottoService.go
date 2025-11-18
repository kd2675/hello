package service

import (
	"lotto/io"
	"lotto/model"
	"lotto/utils"
)

const lottoPrice = 1000

type LottoService interface {
	BuyLottos(wallet *model.Wallet) error
}

type lottoServiceImpl struct{}

func NewLottoService() LottoService {
	return &lottoServiceImpl{}
}

func (s *lottoServiceImpl) BuyLottos(wallet *model.Wallet) error {
	money := wallet.Money()
	count := money / lottoPrice

	for i := 0; i < count; i++ {
		if err := s.buyLotto(wallet); err != nil {
			return err
		}
	}

	io.PrintPurchaseResult(wallet)
	return nil
}

func (s *lottoServiceImpl) buyLotto(wallet *model.Wallet) error {
	lotto := s.generateLotto()
	if lotto == nil {
		return nil
	}

	if err := wallet.Spend(lottoPrice); err != nil {
		return err
	}

	wallet.AddLotto(lotto)
	return nil
}

func (s *lottoServiceImpl) generateLotto() *model.Lotto {
	numbers := utils.PickUniqueNumbersInRange(1, 45, 6)
	lotto, err := model.NewLotto(numbers)
	if err != nil {
		return nil
	}
	return lotto
}

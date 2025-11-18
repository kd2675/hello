package service

import (
	"lotto/io"
	"lotto/model"
)

type WinningService interface {
	WinningLottos(wallet *model.Wallet, winningNumber *model.WinningNumber, bonusNumber *model.BonusNumber)
}

type winningServiceImpl struct{}

func NewWinningService() WinningService {
	return &winningServiceImpl{}
}

func (s *winningServiceImpl) WinningLottos(
	wallet *model.Wallet,
	winningNumber *model.WinningNumber,
	bonusNumber *model.BonusNumber,
) {
	stat := s.checkWinning(wallet, winningNumber, bonusNumber)
	io.PrintWinningStatistics(stat, wallet.SpentMoney())
}

func (s *winningServiceImpl) checkWinning(
	wallet *model.Wallet,
	winningNumber *model.WinningNumber,
	bonusNumber *model.BonusNumber,
) *model.WinningStat {
	stat := model.NewWinningStat()

	for _, lotto := range wallet.Lottos() {
		rank := s.determineRank(lotto, winningNumber, bonusNumber)
		stat.AddRank(rank)
	}

	return stat
}

func (s *winningServiceImpl) determineRank(
	lotto *model.Lotto,
	winningNumber *model.WinningNumber,
	bonusNumber *model.BonusNumber,
) model.Rank {
	matchCount := winningNumber.CountMatch(lotto)
	matchBonus := bonusNumber.Match(lotto)

	return model.GetRank(matchCount, matchBonus)
}

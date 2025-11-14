package model

import "errors"

type Wallet struct {
	money  int
	lottos []*Lotto
}

func NewWallet(money int) (*Wallet, error) {
	if money < 0 {
		return nil, errors.New("금액은 0 이상이어야 합니다")
	}

	return &Wallet{
		money:  money,
		lottos: make([]*Lotto, 0),
	}, nil
}

func (w *Wallet) Money() int {
	return w.money
}

func (w *Wallet) Lottos() []*Lotto {
	return w.lottos
}

func (w *Wallet) Spend(amount int) error {
	if w.money < amount {
		return errors.New("[ERROR] 잔액이 부족합니다")
	}
	w.money -= amount
	return nil
}

func (w *Wallet) AddLotto(lotto *Lotto) {
	w.lottos = append(w.lottos, lotto)
}

func (w *Wallet) LottoCount() int {
	return len(w.lottos)
}

func (w *Wallet) SpentMoney() int {
	return len(w.lottos) * lottoPrice
}

package model

import (
	"errors"
	"strconv"
	"strings"
)

const lottoPrice = 1000

type PurchaseAmount struct {
	price int
}

func NewPurchaseAmount(priceStr string) (*PurchaseAmount, error) {
	price, err := parseAmount(priceStr)
	if err != nil {
		return nil, err
	}

	if err := validateUnit(price); err != nil {
		return nil, err
	}

	return &PurchaseAmount{price: price}, nil
}

func (p *PurchaseAmount) Price() int {
	return p.price
}

func (p *PurchaseAmount) LottoCount() int {
	return p.price / lottoPrice
}

func parseAmount(priceStr string) (int, error) {
	trimmed := strings.TrimSpace(priceStr)
	price, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, errors.New("로또 구입 금액은 숫자여야 합니다")
	}
	return price, nil
}

func validateUnit(price int) error {
	if price%lottoPrice != 0 {
		return errors.New("로또 구입 금액은 1000원 단위 입니다")
	}
	if price <= 0 {
		return errors.New("로또 구입 금액은 양수여야 합니다")
	}
	return nil
}

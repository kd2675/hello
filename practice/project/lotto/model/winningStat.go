package model

type WinningStat struct {
	result map[Rank]int
}

func NewWinningStat() *WinningStat {
	result := make(map[Rank]int)

	for _, rank := range AllWinningRanks() {
		result[rank] = 0
	}

	return &WinningStat{result: result}
}

func (w *WinningStat) AddRank(rank Rank) {
	if rank.IsWinning() {
		w.result[rank]++
	}
}

func (w *WinningStat) GetCount(rank Rank) int {
	return w.result[rank]
}

func (w *WinningStat) CalculateProfitRate(purchaseAmount int) float64 {
	if purchaseAmount == 0 {
		return 0.0
	}
	return float64(w.GetTotalPrize()) / float64(purchaseAmount) * 100.0
}

func (w *WinningStat) GetTotalPrize() int {
	total := 0
	for rank, count := range w.result {
		total += rank.Prize() * count
	}
	return total
}

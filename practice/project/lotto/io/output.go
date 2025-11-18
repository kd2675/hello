package io

import (
	"fmt"
	"lotto/model"
)

func PrintPurchaseResult(wallet *model.Wallet) {
	fmt.Println()
	fmt.Printf("%d개를 구매했습니다.\n", wallet.LottoCount())

	for _, lotto := range wallet.Lottos() {
		fmt.Println(lotto.Numbers())
	}
}

func PrintWinningStatistics(stat *model.WinningStat, purchaseAmount int) {
	printStatisticsHeader()
	printRankStatistics(stat)
	printProfitRate(stat.CalculateProfitRate(purchaseAmount))
}

func printStatisticsHeader() {
	fmt.Println()
	fmt.Println("당첨 통계")
	fmt.Println("---")
}

func printRankStatistics(stat *model.WinningStat) {
	ranks := model.AllWinningRanks()
	for _, rank := range ranks {
		printRank(rank, stat)
	}
}

func printRank(rank model.Rank, stat *model.WinningStat) {
	fmt.Printf("%s (%s원) - %d개\n",
		rank.Description(),
		formatNumber(rank.Prize()),
		stat.GetCount(rank))
}

func printProfitRate(profitRate float64) {
	fmt.Printf("총 수익률은 %.1f%%입니다.\n", profitRate)
}

func PrintErrorMessage(err error) {
	fmt.Printf("[ERROR] %s\n", err.Error())
}

func formatNumber(n int) string {
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	}

	str := fmt.Sprintf("%d", n)
	result := ""
	count := 0

	for i := len(str) - 1; i >= 0; i-- {
		if count > 0 && count%3 == 0 {
			result = "," + result
		}
		result = string(str[i]) + result
		count++
	}

	return result
}

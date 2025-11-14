package model

type Rank int

const (
	None Rank = iota
	Fifth
	Fourth
	Third
	Second
	First
)

var rankInfo = map[Rank]struct {
	matchCount int
	matchBonus bool
	prize      int
	desc       string
}{
	First:  {6, false, 2_000_000_000, "6개 일치"},
	Second: {5, true, 30_000_000, "5개 일치, 보너스 볼 일치"},
	Third:  {5, false, 1_500_000, "5개 일치"},
	Fourth: {4, false, 50_000, "4개 일치"},
	Fifth:  {3, false, 5_000, "3개 일치"},
	None:   {0, false, 0, ""},
}

func GetRank(matchCount int, matchBonus bool) Rank {
	for rank, info := range rankInfo {
		if info.matchCount == matchCount && info.matchBonus == matchBonus {
			return rank
		}
	}
	return None
}

func (r Rank) Prize() int {
	return rankInfo[r].prize
}

func (r Rank) Description() string {
	return rankInfo[r].desc
}

func (r Rank) IsWinning() bool {
	return r != None
}

func AllWinningRanks() []Rank {
	return []Rank{Fifth, Fourth, Third, Second, First}
}

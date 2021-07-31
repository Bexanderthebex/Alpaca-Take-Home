package main

import (
	"strconv"
	"strings"
)

type LotteryBetsVisitor struct {
	boolMap   *BoolMap
	separator string
}

func NewLotteryBetsVisitor(boolMap *BoolMap, separator string) *LotteryBetsVisitor {
	return &LotteryBetsVisitor{
		boolMap:   boolMap,
		separator: separator,
	}
}

func (l *LotteryBetsVisitor) Visit(stringifiedLottoBets *string) {
	lottoBets := strings.Split(*stringifiedLottoBets, l.separator)

	for _, lottoBet := range lottoBets {
		lottoPicks := make([]string, 0)
		for _, s := range strings.Split(lottoBet, " ") {
			if s != "" {
				lottoPicks = append(lottoPicks, s)
			}
		}

		formattedLottoPicks := make([]uint, 0)
		for _, lottoPick := range lottoPicks {
			formattedLottoPick, _ := strconv.Atoi(lottoPick)
			validLottoPickFormat := uint(formattedLottoPick)
			formattedLottoPicks = append(formattedLottoPicks, validLottoPickFormat)
		}

		if picksValid(formattedLottoPicks) {
			for _, flt := range formattedLottoPicks {
				recordId := l.boolMap.GetTotalRecords()
				l.boolMap.SetValue(flt, recordId, true)
			}

			l.boolMap.IncrementTotalRecords()
		}

	}
}

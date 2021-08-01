package main

import (
	"fmt"
	"os"
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

func (l *LotteryBetsVisitor) Visit(lottoBet string) {
	// TODO: made the cap similar to width
	lottoPicks := make([]string, 0, 5)
	for _, s := range strings.Split(lottoBet, l.separator) {
		if s != "" {
			lottoPicks = append(lottoPicks, s)
		}
	}

	formattedLottoPicks := make([]uint, 0, 5)
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

	if !picksValid(formattedLottoPicks) && os.Getenv("APP_ENVIRONMENT") != "TEST" {
		fmt.Println("Detected invalid lotto picks. Discarding")
		fmt.Println(lottoPicks)
	}
}

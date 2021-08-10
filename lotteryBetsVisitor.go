package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LotteryBetsVisitor struct {
	bitmap    BitMapIndex
	separator string
}

func NewLotteryBetsVisitor(bitmap BitMapIndex, separator string) *LotteryBetsVisitor {
	return &LotteryBetsVisitor{
		bitmap:    bitmap,
		separator: separator,
	}
}

func (l *LotteryBetsVisitor) Visit(lottoBet string) {
	// TODO: made the cap similar to width
	lottoPicks := make([]string, 0, 5)
	for _, s := range strings.Split(lottoBet, l.separator) {
		lottoPicks = append(lottoPicks, s)
	}

	formattedLottoPicks := make(map[uint]uint)
	for _, lottoPick := range lottoPicks {
		validLottoPickFormat, _ := strconv.Atoi(lottoPick)
		formattedLottoPicks[uint(validLottoPickFormat)] += 1
	}

	if picksValid(formattedLottoPicks) {
		recordId := l.bitmap.GetTotalRecords()
		for flt := range formattedLottoPicks {
			l.bitmap.SetValue(flt, recordId, true)
		}

		l.bitmap.IncrementTotalRecords()
	}

	if !picksValid(formattedLottoPicks) && os.Getenv("APP_ENVIRONMENT") != "TEST" {
		fmt.Println("Detected invalid lotto picks. Discarding")
		fmt.Println(lottoPicks)
	}
}

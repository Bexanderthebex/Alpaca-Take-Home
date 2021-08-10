package main

import (
	"testing"
)

func TestLotteryBetsQueryEngine_BitMap_ExecuteQuery(t *testing.T) {
	bitmap := NewBitMap(minimumValidPick, maximumValidPick, 100)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitmap, " ")

	mockBets := []string{
		"29 10 11 12 13",
		"20 31 32 81 60",
		"78 30 29 10 11",
		"29 30 32 31 80",
		"29 30 32 31 78",
	}

	expectedResult := make(map[uint]uint)
	expectedResult[5] = 1
	expectedResult[4] = 1
	expectedResult[3] = 1
	expectedResult[2] = 1

	for _, mockBet := range mockBets {
		lotteryBetsVisitor.Visit(mockBet)
	}

	winningPicks := []uint{29, 30, 32, 31, 78}
	queryPlan := NewQueryPlan(SELECT, true, bitmap)
	queryPlan.SetColumnsToSelect(&winningPicks)
	queryPlan.SetAggregationStrategy(NewQueryAggregation(bitmap.GetTotalRecords()))
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)

	queryEngine := LotteryBetsQueryEngine{
		bitmap: bitmap,
	}

	answerMap := queryEngine.ExecuteQuery(queryPlan)

	assertEqualMap(t, expectedResult, answerMap, "Wrong query result")
}

func BenchmarkLotteryBetsQueryEngine_BitMap_Constant_ExecuteQuery(b *testing.B) {
	bitMap := NewBitMap(minimumValidPick, maximumValidPick, maximumBettors/8)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	lottoBet := "29 32 34 78 39"
	for i := 0; i < maximumBettors; i++ {
		lotteryBetsVisitor.Visit(lottoBet)
	}

	winningPicks := []uint{29, 32, 34, 78, 39}
	queryPlan := NewQueryPlan(SELECT, true, bitMap)
	queryPlan.SetColumnsToSelect(&winningPicks)
	queryPlan.SetAggregationStrategy(NewQueryAggregation(bitMap.GetTotalRecords()))
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)

	queryEngine := LotteryBetsQueryEngine{
		bitmap: bitMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

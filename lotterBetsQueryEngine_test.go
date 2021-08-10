package main

import (
	"testing"
)

func TestLotteryBetsQueryEngine_BoolMap_ExecuteQuery(t *testing.T) {
	boolMap := NewBoolMap(minimumValidPick, maximumValidPick, 100)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

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

	winningPicks := make(map[uint]uint)
	winningPicks[29] = 1
	winningPicks[30] = 1
	winningPicks[32] = 1
	winningPicks[31] = 1
	winningPicks[78] = 1
	accumulator := make([]uint, 5)
	queryPlan := NewSelectQueryPlan(boolMap)
	queryPlan.SetColumnsToSelect(&winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(&winningPicks, &accumulator, boolMap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: boolMap,
	}

	answerMap := queryEngine.ExecuteQuery(queryPlan)

	assertEqualMap(t, expectedResult, answerMap, "Wrong query result")
}

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

	winningPicks := make(map[uint]uint)
	winningPicks[29] = 1
	winningPicks[30] = 1
	winningPicks[32] = 1
	winningPicks[31] = 1
	winningPicks[78] = 1
	accumulator := make([]uint, 5)
	queryPlan := NewSelectQueryPlan(bitmap)
	queryPlan.SetColumnsToSelect(&winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(&winningPicks, &accumulator, bitmap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitmap,
	}

	answerMap := queryEngine.ExecuteQuery(queryPlan)

	assertEqualMap(t, expectedResult, answerMap, "Wrong query result")
}

func BenchmarkLotteryBetsQueryEngine_BoolMap_Constant_ExecuteQuery(b *testing.B) {
	boolMap := NewBoolMap(minimumValidPick, maximumValidPick, maximumBettors+1)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

	lottoBet := "29 32 34 78 39"
	for i := 0; i < maximumBettors; i++ {
		lotteryBetsVisitor.Visit(lottoBet)
	}

	winningPicks := make(map[uint]uint)
	winningPicks[29] = 1
	winningPicks[32] = 1
	winningPicks[34] = 1
	winningPicks[78] = 1
	winningPicks[39] = 1
	accumulator := make([]uint, maximumBettors)
	queryPlan := NewSelectQueryPlan(boolMap)
	queryPlan.SetColumnsToSelect(&winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(&winningPicks, &accumulator, boolMap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: boolMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

func BenchmarkLotteryBetsQueryEngine_BitMap_Constant_ExecuteQuery(b *testing.B) {
	bitMap := NewBitMap(minimumValidPick, maximumValidPick, maximumBettors)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	lottoBet := "29 32 34 78 39"
	for i := 0; i < maximumBettors; i++ {
		lotteryBetsVisitor.Visit(lottoBet)
	}

	winningPicks := make(map[uint]uint)
	winningPicks[29] = 1
	winningPicks[32] = 1
	winningPicks[34] = 1
	winningPicks[78] = 1
	winningPicks[39] = 1
	accumulator := make([]uint, maximumBettors)
	queryPlan := NewSelectQueryPlan(bitMap)
	queryPlan.SetColumnsToSelect(&winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(&winningPicks, &accumulator, bitMap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

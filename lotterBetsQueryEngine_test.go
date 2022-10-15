package main

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLotteryBetsQueryEngine_ExecuteQuery_Unoptimized_Happy(t *testing.T) {
	bitmap := NewBitMap(minimumValidPick, maximumValidPick, 100)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitmap, " ")

	expectedResult := make(map[uint]uint)
	expectedResult[5] = 1
	expectedResult[4] = 1
	expectedResult[3] = 1
	expectedResult[2] = 1

	for _, mockBet := range mockBets {
		lotteryBetsVisitor.Visit(mockBet)
	}

	accumulator := make([]uint, 5)
	queryPlan := NewSelectQueryPlan(bitmap)
	queryPlan.SetColumnsToSelect(mockPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(mockPicks, &accumulator, bitmap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitmap,
	}

	answerMap := queryEngine.ExecuteQuery(queryPlan)

	assert.Equal(t, expectedResult, answerMap, "Wrong query result")
}

func TestLotteryBetsQueryEngine_ExecuteQuery_Unoptimized_Unhappy(t *testing.T) {
	bitmap := NewBitMap(minimumValidPick, maximumValidPick, 100)

	// Don't instantiate the bitmap by skipping the data loading part
	expectedResult := make(map[uint]uint)
	expectedResult[5] = 1
	expectedResult[4] = 1
	expectedResult[3] = 1
	expectedResult[2] = 1

	accumulator := make([]uint, 5)
	queryPlan := NewSelectQueryPlan(bitmap)
	queryPlan.SetColumnsToSelect(mockPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(mockPicks, &accumulator, bitmap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitmap,
	}

	answerMap := queryEngine.ExecuteQuery(queryPlan)

	assert.NotEqual(t, expectedResult, answerMap, "Query result is expected to not match")
}

func TestLotteryBetsQueryEngine_ExecuteQuery_Optimized_Happy(t *testing.T) {
	bitmap := NewBitMap(minimumValidPick, maximumValidPick, 100)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitmap, " ")

	expectedResult := make(map[uint]uint)
	expectedResult[5] = 1
	expectedResult[4] = 1
	expectedResult[3] = 1
	expectedResult[2] = 1

	for _, mockBet := range mockBets {
		lotteryBetsVisitor.Visit(mockBet)
	}

	queryPlan := NewSelectQueryPlan(bitmap)
	queryPlan.SetColumnsToSelect(mockPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewOptimizedAggregation(bitmap, mockPicks, 2, 5))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitmap,
	}

	answerMap := queryEngine.ExecuteQuery(queryPlan)

	assert.Equal(t, expectedResult, answerMap, "Wrong query result")
}

func TestLotteryBetsQueryEngine_ExecuteQuery_Optimized_Unhappy(t *testing.T) {
	bitmap := NewBitMap(minimumValidPick, maximumValidPick, 100)

	// Don't instantiate the bitmap by skipping the data loading part
	expectedResult := make(map[uint]uint)
	expectedResult[5] = 1
	expectedResult[4] = 1
	expectedResult[3] = 1
	expectedResult[2] = 1

	queryPlan := NewSelectQueryPlan(bitmap)
	queryPlan.SetColumnsToSelect(mockPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewOptimizedAggregation(bitmap, mockPicks, 2, 5))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitmap,
	}

	answerMap := queryEngine.ExecuteQuery(queryPlan)

	assert.NotEqual(t, expectedResult, answerMap, "Query result is expected to not match")
}

func BenchmarkLotteryBetsQueryEngine_ExecuteQuery_Unoptimized_10m_Big_Theta(b *testing.B) {
	bitMap := NewBitMap(minimumValidPick, maximumValidPick, maximumBettors)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	file, err := os.Open("10m-v2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lottoBet := scanner.Text()
		lotteryBetsVisitor.Visit(lottoBet)
	}

	file.Close()

	winningPicks := make(map[uint]uint)
	winningPicks[29] = 1
	winningPicks[32] = 1
	winningPicks[34] = 1
	winningPicks[78] = 1
	winningPicks[39] = 1
	accumulator := make([]uint, maximumBettors)
	queryPlan := NewSelectQueryPlan(bitMap)
	queryPlan.SetColumnsToSelect(winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(mockPicks, &accumulator, bitMap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

func BenchmarkLotteryBetsQueryEngine_ExecuteQuery_Unoptimized_10m_Big_Oh(b *testing.B) {
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
	queryPlan.SetColumnsToSelect(winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(mockPicks, &accumulator, bitMap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

func BenchmarkLotteryBetsQueryEngine_ExecuteQuery_Unoptimized_10m_Big_Omega(b *testing.B) {
	bitMap := NewBitMap(minimumValidPick, maximumValidPick, maximumBettors)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	lottoBet := "29 32 34 78 39"
	for i := 0; i < maximumBettors; i++ {
		lotteryBetsVisitor.Visit(lottoBet)
	}

	winningPicks := make(map[uint]uint)
	winningPicks[1] = 1
	winningPicks[2] = 1
	winningPicks[3] = 1
	winningPicks[4] = 1
	winningPicks[5] = 1
	accumulator := make([]uint, maximumBettors)
	queryPlan := NewSelectQueryPlan(bitMap)
	queryPlan.SetColumnsToSelect(winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewCountAggregation(mockPicks, &accumulator, bitMap))
	queryPlan.AddAggregationStrategy(NewGroupAggregation(&accumulator, []uint{5, 4, 3, 2}))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

func BenchmarkLotteryBetsQueryEngine_ExecuteQuery_Optimized_10m_Big_Theta(b *testing.B) {
	bitMap := NewBitMap(minimumValidPick, maximumValidPick, maximumBettors)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	file, err := os.Open("10m-v2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lottoBet := scanner.Text()
		lotteryBetsVisitor.Visit(lottoBet)
	}

	file.Close()

	winningPicks := make(map[uint]uint)
	winningPicks[29] = 1
	winningPicks[32] = 1
	winningPicks[34] = 1
	winningPicks[78] = 1
	winningPicks[39] = 1
	queryPlan := NewSelectQueryPlan(bitMap)
	queryPlan.SetColumnsToSelect(winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewOptimizedAggregation(bitMap, winningPicks, 2, 5))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

func BenchmarkLotteryBetsQueryEngine_ExecuteQuery_Optimized_10m_Big_Oh(b *testing.B) {
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
	queryPlan := NewSelectQueryPlan(bitMap)
	queryPlan.SetColumnsToSelect(winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewOptimizedAggregation(bitMap, winningPicks, 2, 5))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

func BenchmarkLotteryBetsQueryEngine_ExecuteQuery_Optimized_10m_Big_Omega(b *testing.B) {
	bitMap := NewBitMap(minimumValidPick, maximumValidPick, maximumBettors)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	lottoBet := "29 32 34 78 39"
	for i := 0; i < maximumBettors; i++ {
		lotteryBetsVisitor.Visit(lottoBet)
	}

	winningPicks := make(map[uint]uint)
	winningPicks[1] = 1
	winningPicks[2] = 1
	winningPicks[3] = 1
	winningPicks[4] = 1
	winningPicks[5] = 1
	queryPlan := NewSelectQueryPlan(bitMap)
	queryPlan.SetColumnsToSelect(winningPicks)
	queryPlan.SetMinValue(2)
	queryPlan.SetMaxValue(5)
	queryPlan.AddAggregationStrategy(NewOptimizedAggregation(bitMap, winningPicks, 2, 5))

	queryEngine := LotteryBetsQueryEngine{
		bitmapIndex: bitMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

package main

import "testing"

func TestLotteryBetsQueryEngine_ExecuteQuery(t *testing.T) {
	boolMap := NewBoolMap(1, 90, 5)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

	mockBets := []string{
		"29 10 11 12 13",
		"20 31 32 81 60",
		"78 30 29 10 11",
		"29 30 32 31 80",
		"29 30 32 31 78",
	}

	for _, mockBet := range mockBets {
		lotteryBetsVisitor.Visit(mockBet)
	}

	expectedResult := make(map[uint]uint)
	expectedResult[5] = 1
	expectedResult[4] = 1
	expectedResult[3] = 1
	expectedResult[2] = 1

	winningPicks := []uint{29, 30, 32, 31, 78}
	queryPlan := QueryPlan{
		columnsToSelect: &winningPicks,
		aggregationCmd:  NewQueryAggregationBool(boolMap.GetTotalRecords()),
		minValue:        minimumValidPick,
		maxValue:        maximumValidPick,
		category:        true,
	}

	queryEngine := LotteryBetsQueryEngine{
		boolMap: boolMap,
	}

	answerMap := queryEngine.ExecuteQuery(queryPlan)

	assertEqualMap(t, expectedResult, answerMap, "Wrong query result")
}

func BenchmarkLotteryBetsQueryEngine_ExecuteQuery(b *testing.B) {
	boolMap := NewBoolMap(1, 90, b.N)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

	bet := "29 32 34 78 39"
	for i := 0; i < b.N; i++ {
		lotteryBetsVisitor.Visit(bet)
	}

	winningPicks := []uint{29, 32, 34, 78, 39}
	queryPlan := QueryPlan{
		columnsToSelect: &winningPicks,
		aggregationCmd:  NewQueryAggregationBool(boolMap.GetTotalRecords()),
		minValue:        minimumValidPick,
		maxValue:        maximumValidPick,
		category:        true,
	}

	queryEngine := LotteryBetsQueryEngine{
		boolMap: boolMap,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queryEngine.ExecuteQuery(queryPlan)
	}
}

package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestLotteryBetsQueryEngine_ExecuteQuery(t *testing.T) {
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

func BenchmarkLotteryBetsQueryEngine_BoolMap_10m_v2_ExecuteQuery(b *testing.B) {
	// Initialize data first, so need to stop the timer
	b.StopTimer()
	file, err := os.Open("10m-v2.txt")
	if err != nil {
		log.Fatal(err)
	}

	boolMap := NewBoolMap(minimumValidPick, maximumValidPick, maximumBettors)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lottoBet := scanner.Text()
		lotteryBetsVisitor.Visit(lottoBet)
	}
	file.Close()

	var winningPicks *[]uint
	queryPlan := QueryPlan{
		columnsToSelect: winningPicks,
		aggregationCmd:  NewQueryAggregationBool(boolMap.GetTotalRecords()),
		minValue:        minimumValidPick,
		maxValue:        maximumValidPick,
		category:        true,
	}

	queryEngine := LotteryBetsQueryEngine{
		boolMap: boolMap,
	}

	for i := 0; i < b.N; i++ {
		// stop the timer again when generating a new pick
		b.StopTimer()
		newWinningPick := make([]uint, 0)
		for j := 0; j < 5; j++ {
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 90
			newWinningPick = append(newWinningPick, uint(rand.Intn(max-min+1)+min))
		}
		winningPicks = &newWinningPick
		queryPlan.columnsToSelect = winningPicks
		b.StartTimer()
		queryEngine.ExecuteQuery(queryPlan)
	}
}

// go test -bench=bitMapExecuteQuery
func BenchmarkLotteryBetsQueryEngine_BitMap_10m_v2_ExecuteQuery(b *testing.B) {
	// Initialize data first, so need to stop the timer
	b.StopTimer()
	file, err := os.Open("10m-v2.txt")
	if err != nil {
		log.Fatal(err)
	}

	bitMap := NewBitMap(minimumValidPick, maximumValidPick, maximumBettors/8)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lottoBet := scanner.Text()
		lotteryBetsVisitor.Visit(lottoBet)
	}
	file.Close()

	var winningPicks *[]uint
	queryPlan := QueryPlan{
		columnsToSelect: winningPicks,
		aggregationCmd:  NewQueryAggregationBool(bitMap.GetTotalRecords()),
		minValue:        minimumValidPick,
		maxValue:        maximumValidPick,
		category:        true,
	}

	queryEngine := LotteryBetsQueryEngine{
		boolMap: bitMap,
	}

	for i := 0; i < b.N; i++ {
		// stop the timer again when generating a new pick
		b.StopTimer()
		newWinningPick := make([]uint, 0)
		for j := 0; j < 5; j++ {
			rand.Seed(time.Now().UnixNano())
			min := 1
			max := 90
			newWinningPick = append(newWinningPick, uint(rand.Intn(max-min+1)+min))
		}
		winningPicks = &newWinningPick
		queryPlan.columnsToSelect = winningPicks
		b.StartTimer()
		queryEngine.ExecuteQuery(queryPlan)
	}
}

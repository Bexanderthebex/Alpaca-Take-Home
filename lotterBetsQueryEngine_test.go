package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func BenchmarkLotteryBetsQueryEngine_10m_v2_ExecuteQuery(b *testing.B) {
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

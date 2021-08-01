package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// valid value range for lottery picks
var minimumValidPick uint = 1
var maximumValidPick uint = 90

var lottoPickLength = 5

var maximumBettors = 10000000

func main() {
	// call file
	fileToIndex := os.Args[1:2]

	file, err := os.Open(fileToIndex[0])
	if err != nil {
		log.Fatal(err)
	}

	//playerPicksIndex := lib.New(minimumValidPick, maximumValidPick, 10000000)

	boolMap := NewBoolMap(minimumValidPick, maximumValidPick, maximumBettors)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

	for {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lottoBet := scanner.Text()
			lotteryBetsVisitor.Visit(lottoBet)
		}

		break
	}

	if fileCloseError := file.Close(); fileCloseError != nil {
		panic(fileCloseError)
	}

	fmt.Println("READY")

	// query engine
	for continueSearch := true; continueSearch == true; {
		var winningPick1 uint
		var winningPick2 uint
		var winningPick3 uint
		var winningPick4 uint
		var winningPick5 uint
		noOfWinningPicksParsed, winningPickParsingError := fmt.Scanf("%d %d %d %d %d\n", &winningPick1, &winningPick2, &winningPick3, &winningPick4, &winningPick5)
		if winningPickParsingError != nil {
			fmt.Println(winningPickParsingError)
			fmt.Println("Please enter the correct format")
		}
		if noOfWinningPicksParsed < 5 {
			continueSearch = false
			fmt.Printf("Only %d\n", noOfWinningPicksParsed)
		}

		fmt.Printf("Sucessfully parsed %d\n winning picks", noOfWinningPicksParsed)

		winningPicks := []uint{winningPick1, winningPick2, winningPick3, winningPick4, winningPick5}
		isPicksValid := picksValid(winningPicks)
		if !isPicksValid {
			fmt.Println("Please enter values that are only between 1 and 90")
			continue
		}

		fmt.Println(winningPicks)

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

		//answersMap := calculateWinners(winningPicks, playerPicksIndex, recordCount)
		answersMap := queryEngine.ExecuteQuery(queryPlan)

		for i := 5; i >= 2; i-- {
			fmt.Printf("%d: %d\n", i, answersMap[uint(i)])
		}
	}
}

func picksValid(picks []uint) bool {
	if len(picks) < lottoPickLength {
		return false
	}

	picksMap := make(map[uint]int)

	for _, v := range picks {
		if !isNumberValid(int(v)) {
			return false
		}

		picksMap[v] += 1
	}

	for _, v := range picksMap {
		if v > 1 {
			return false
		}
	}

	return true
}

func isNumberValid(val int) bool {
	if val >= int(minimumValidPick) && val <= int(maximumValidPick) {
		return true
	}

	return false
}

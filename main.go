package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	bitMap := NewBitMap(minimumValidPick, maximumValidPick, maximumBettors/8)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lottoBet := scanner.Text()
		lotteryBetsVisitor.Visit(lottoBet)
	}

	if fileCloseError := file.Close(); fileCloseError != nil {
		panic(fileCloseError)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("READY")

	// query engine
	for continueSearch := true; continueSearch == true; {
		text, readStringError := reader.ReadString('\n')
		if readStringError != nil {
			fmt.Println(readStringError)
			fmt.Println("Please enter the correct format")
		}

		text = strings.Replace(text, "\n", "", -1)
		winningPicks := make([]uint, 0)
		for _, bet := range strings.Split(text, " ") {
			if bet != "" {
				formattedBetString, _ := strconv.Atoi(bet)
				winningPicks = append(winningPicks, uint(formattedBetString))
			}
		}

		if len(winningPicks) < 5 {
			fmt.Printf("Please enter 5 winning picks, only %d were entered\n", len(winningPicks))
			continue
		}

		isPicksValid := picksValid(winningPicks)
		if !isPicksValid {
			fmt.Println("Please enter values that are only between 1 and 90")
			continue
		}

		fmt.Println("Winning picks parsed:")
		fmt.Println(winningPicks)

		queryPlan := QueryPlan{
			columnsToSelect: &winningPicks,
			aggregationCmd:  NewQueryAggregationBool(bitMap.GetTotalRecords()),
			minValue:        minimumValidPick,
			maxValue:        maximumValidPick,
			category:        true,
		}

		queryEngine := LotteryBetsQueryEngine{
			boolMap: bitMap,
		}

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

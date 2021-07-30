package main

import (
	"alpacahq-take-home/m/lib"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var minimumValidPick uint = 1
var maximumValidPick uint = 90

func main() {
	// call file
	fileToIndex := os.Args[1:2]

	file, err := os.Open(fileToIndex[0])
	if err != nil {
		log.Fatal(err)
	}

	playerPicksIndex := lib.New(minimumValidPick, maximumValidPick, 10000000)

	var recordCount uint
	for {
		var lottoPick1 uint
		var lottoPick2 uint
		var lottoPick3 uint
		var lottoPick4 uint
		var lottoPick5 uint
		_, err2 := fmt.Fscanf(file, "%d %d %d %d %d\n", &lottoPick1, &lottoPick2, &lottoPick3, &lottoPick4, &lottoPick5)
		if errors.Is(err2, io.EOF) {
			break
		}
		picks := []uint{lottoPick1, lottoPick2, lottoPick3, lottoPick4, lottoPick5}

		isLottoPickValid := picksValid(picks)

		// skip adding it to the record if it is not valid
		if !isLottoPickValid {
			continue
		}

		for i := 0; i < len(picks); i++ {
			pick := picks[i]
			playerPicksIndex.BitWiseOr(pick, recordCount)
		}
		recordCount += 1
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

		answersMap := calculateWinners(winningPicks, playerPicksIndex, recordCount)

		for i := 5; i >= 2; i-- {
			fmt.Printf("%d: %d\n", i, answersMap[uint(i)])
		}
	}
}

func calculateWinners(winningPicks []uint, playerPicksIndex *lib.BitMap, totalBettors uint) map[uint]int {
	winningPicksMiniIndex := sync.Map{}

	var waitgroup sync.WaitGroup
	waitgroup.Add(len(winningPicks))

	for _, winningPick := range winningPicks {
		go convertByteIndexToBool(playerPicksIndex, &waitgroup, winningPick, &winningPicksMiniIndex)
	}

	waitgroup.Wait()

	answer := make(map[uint]int)
	for i := 0; uint(i) < totalBettors; i++ {
		var noOfHits uint
		winningPicksMiniIndex.Range(func(syncMapKey interface{}, syncMapValue interface{}) bool {
			boolIndex := syncMapValue.(*[]bool)
			if (*boolIndex)[i] {
				noOfHits += 1
			}

			return true
		})
		if noOfHits >= 2 {
			answer[noOfHits] += 1
		}
	}

	return answer
}

func convertByteIndexToBool(bitMap *lib.BitMap, waitgroup *sync.WaitGroup, index uint, memoryIndex *sync.Map) {
	defer waitgroup.Done()
	boolIndex := bitMap.GetIndexAsBool(index)
	memoryIndex.Store(index, boolIndex)
}

func picksValid(picks []uint) bool {
	picksMap := make(map[uint]int)

	for _, v := range picks {
		if !isNumberValid(v) {
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

func isNumberValid(val uint) bool {
	if val >= minimumValidPick && val <= maximumValidPick {
		return true
	}

	return false
}

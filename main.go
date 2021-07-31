package main

import (
	"alpacahq-take-home/m/lib"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// valid value range for lottery picks
var minimumValidPick uint = 1
var maximumValidPick uint = 90

var lottoPickLength = 5

type BitSet map[uint]*[]bool

// totalRecords could be a BitSet value
func (bs *BitSet) CalculateWinners(nums *[]uint, minMatch uint, totalRecords uint) map[uint]uint {
	winningPicksIndex := make(map[uint]*[]bool)
	// winning picks
	for _, v := range *nums {
		winningPicksIndex[v] = (*bs)[v]
	}

	groupedWinnersCount := make(map[uint]uint)
	for i := uint(len(*nums)); i >= 2; i-- {
		groupedWinnersCount[i] = 0
	}

	for i := uint(0); i < totalRecords; i++ {
		var noOfHits uint = 0
		for _, index := range winningPicksIndex {
			if (*index)[i] {
				noOfHits += 1
			}
		}

		if noOfHits >= minMatch {
			groupedWinnersCount[noOfHits] += 1
		}
	}

	return groupedWinnersCount
}

func main() {
	// call file
	fileToIndex := os.Args[1:2]

	file, err := os.Open(fileToIndex[0])
	if err != nil {
		log.Fatal(err)
	}

	//playerPicksIndex := lib.New(minimumValidPick, maximumValidPick, 10000000)

	bitSet := make(BitSet)
	for i := minimumValidPick; i <= maximumValidPick; i++ {
		boolSet := make([]bool, 10000000)
		bitSet[i] = &boolSet
	}

	var recordCount uint = 0
	var bufferOffset int64 = 0

	for {
		fileReadBuffer := make([]byte, 4096*3)
		_, fileReadError := file.ReadAt(fileReadBuffer, bufferOffset)
		if errors.Is(fileReadError, io.EOF) {
			break
		}

		// find the last complete line scanned
		for i := len(fileReadBuffer) - 1; i > 0; i -= 1 {
			if string(fileReadBuffer[i]) == "\n" {
				fileReadBuffer = fileReadBuffer[:i]
				break
			}
		}
		bufferOffset += int64(len(fileReadBuffer) + 1)

		stringifiedLottoRecords := string(fileReadBuffer)
		records := strings.Split(stringifiedLottoRecords, "\n")
		for _, record := range records {
			lottoPicks := make([]string, 0)
			for _, s := range strings.Split(record, " ") {
				if s != "" {
					lottoPicks = append(lottoPicks, s)
				}
			}

			formattedLottoPicks := make([]uint, 0)
			for _, lottoPick := range lottoPicks {
				formattedLottoPick, _ := strconv.Atoi(lottoPick)
				validLottoPickFormat := uint(formattedLottoPick)
				formattedLottoPicks = append(formattedLottoPicks, validLottoPickFormat)
			}

			if picksValid(formattedLottoPicks) {
				for _, flt := range formattedLottoPicks {
					boolMap := bitSet[flt]
					(*boolMap)[recordCount] = true
				}

				recordCount += 1
			}

		}
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

		//answersMap := calculateWinners(winningPicks, playerPicksIndex, recordCount)
		answersMap := bitSet.CalculateWinners(&winningPicks, 2, recordCount)

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

package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"sync"
)

// Think of this as a block of memory to which an index points to
// Records UUIDs of lottery picks made by a player
// My mental model of a player entry into the the game
type PlayerPicksIndex struct {
	IDs					map[string]bool
}

func main() {
	fileToIndex := os.Args[1:2]

	file, err := os.Open(fileToIndex[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// This serves as the lookup table for occurrence of picks based on an existence of a number match in the player's pick
	// Space Complexity O(n)
	playerPickIndex := make(map[int]*PlayerPicksIndex)

	// One improvement I would make to make the loading faster is to make a channel and create workers that would build the index
	// However, that does not come without a trade off. The memory requirement to load the data will increase since one will need Sync Map for that
	// The solution already uses a UUID to make it flexible for that change.

	// Also, since load time is not that much of a concern from the problem, I decided not to optimize around that

	// I am also sure that there is something that can be done with the size of the data — e.g. compacting the data size
	// However, I haven't done that in the past but if given enough time to research, I can definitely deliver that
	for  {
		var lottoPick1 int
		var lottoPick2 int
		var lottoPick3 int
		var lottoPick4 int
		var lottoPick5 int
		_, err2 := fmt.Fscanf(file, "%d %d %d %d %d\n", &lottoPick1, &lottoPick2, &lottoPick3, &lottoPick4, &lottoPick5)
		if errors.Is(err2, io.EOF) {
			break
		}
		picks := []int{lottoPick1, lottoPick2, lottoPick3, lottoPick4, lottoPick5 }

		// The unique identifier tied to a string
		recordId := uuid.NewString()
		// Step 1: Iterate through the picks of the player
		for i := 0 ; i < len(picks) ; i ++ {
			pick := picks[i]
			// Step 2: Initialize the memory block for a number match if it does not exists
			// Otherwise, just store the ID regardless
			if _, exists := playerPickIndex[pick]; !exists {
				playerPickIndex[pick] = &PlayerPicksIndex{IDs: make(map[string]bool)}
			}
			playerPickIndex[pick].IDs[recordId] = true
		}
	}

	fmt.Println("READY")

	for continueSearch := true ; continueSearch == true ; {
		var winningPick1 int
		var winningPick2 int
		var winningPick3 int
		var winningPick4 int
		var winningPick5 int
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

		winningPicks := []int{winningPick1, winningPick2, winningPick3, winningPick4, winningPick5 }
		fmt.Println(winningPicks)

		// Do the searching here
		answersMap := make(map[string]int)
		// I thought of the number of winners per match as a bucket that are independent of each other
		// Therefore, I used a sync.Map to be able to do concurrent writes
		groupedWinners := sync.Map{}
		for i := 2 ; i <= 5 ; i ++ {
			groupedWinners.Store(i, 0)
		}

		// STEP 1: Iterate through the winning picks entered by the user
		// O(x) x 1-5
		for i := 0; i < len(winningPicks) ; i ++ {
			winningPick := winningPicks[i]

			// STEP 2: Fetch the index block through the lookup table that contains the IDs of the lottery picks that matched
			winningPickIndexBlock := playerPickIndex[winningPick].IDs
			// Time Complexity O(n)
			for k, _ := range winningPickIndexBlock {
				// STEP 3: Dedupe and store the number of instances the IDs matched a winning pick
				if _, exists := answersMap[k]; !exists {
					answersMap[k] = 0
				}
				answersMap[k] += 1
			}
		}

		// STEP 1: Initialize concurrent read and write of the grouped answers
		var wg sync.WaitGroup
		wg.Add(4)

		// STEP 2: Compartmentalize and concurrently search for winners per no of matched picks in the winning list of picks
		for i := 2; i <= 5; i++ {
			go findAnswers(&answersMap, &groupedWinners, &wg, i)
		}

		wg.Wait()

		// Print the answers
		for i := 5 ; i >= 2 ; i -- {
			noOfHits, _ := groupedWinners.Load(i)
			fmt.Printf("%d: %d\n", i, noOfHits.(int))
		}
	}
}

func findAnswers(answersMap *map[string]int, ans *sync.Map, waitgroup *sync.WaitGroup, mapKey int) {
	defer waitgroup.Done()
	for _, v := range *(answersMap) {
		if v == mapKey {
			occurenceCount, _ := ans.Load(v)
			ans.Store(v, occurenceCount.(int) + 1)
		}
	}
}

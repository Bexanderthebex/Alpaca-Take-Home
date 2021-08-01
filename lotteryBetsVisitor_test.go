package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

//TODO: not bad but might be improved
func BenchmarkLotteryBetsVisitor_Visit(b *testing.B) {
	boolMap := NewBoolMap(1, 90, 10000000)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

	file, err := os.Open("file-mocks/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lottoBet := scanner.Text()
			b.StartTimer()
			lotteryBetsVisitor.Visit(lottoBet)
			b.StopTimer()
		}
		break
	}
}

func TestLotteryBetsVisitor_Visit(t *testing.T) {
	boolMap := NewBoolMap(1, 90, 100)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

	mockLotteryBets := []string{
		"29 30 31 33 90",
		"19 60 61 23 89",
	}

	for _, v := range mockLotteryBets {
		lotteryBetsVisitor.Visit(v)
	}

	assertEqualInt(t, 2, int(boolMap.GetTotalRecords()), "Total should be 2")
	// Test bet 1
	assertEqualBool(t, true, boolMap.GetValue(29, 0), "Value should be true")
	assertEqualBool(t, true, boolMap.GetValue(30, 0), "Value should be true")
	assertEqualBool(t, true, boolMap.GetValue(31, 0), "Value should be true")
	assertEqualBool(t, true, boolMap.GetValue(33, 0), "Value should be true")
	assertEqualBool(t, true, boolMap.GetValue(90, 0), "Value should be true")
	// Test bet 2
	assertEqualBool(t, true, boolMap.GetValue(19, 1), "Value should be true")
	assertEqualBool(t, true, boolMap.GetValue(60, 1), "Value should be true")
	assertEqualBool(t, true, boolMap.GetValue(61, 1), "Value should be true")
	assertEqualBool(t, true, boolMap.GetValue(23, 1), "Value should be true")
	assertEqualBool(t, true, boolMap.GetValue(89, 1), "Value should be true")
	// Falsy tests
	assertEqualBool(t, false, boolMap.GetValue(18, 1), "Value should be true")
	assertEqualBool(t, false, boolMap.GetValue(28, 0), "Value should be true")
}

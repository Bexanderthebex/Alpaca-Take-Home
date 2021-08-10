package main

import (
	"testing"
)

func BenchmarkBoolMap_SetValue(b *testing.B) {
	boolMap := NewBoolMap(1, 90, b.N)

	for i := 0; i < b.N; i++ {
		boolMap.SetValue(uint(90), uint(i), true)
	}
}

func BenchmarkBoolMap_GetValue(b *testing.B) {
	boolMap := NewBoolMap(1, 90, b.N)

	for i := 0; i < b.N; i++ {
		boolMap.GetValue(uint(90), uint(i))
	}
}

func BenchmarkBoolMap_CalculateWinners(b *testing.B) {
	boolMap := NewBoolMap(1, 90, b.N)
	lotteryBetsVisitor := NewLotteryBetsVisitor(boolMap, " ")

	bet := "29 32 34 78 39"
	for i := 0; i < b.N; i++ {
		lotteryBetsVisitor.Visit(bet)
	}

	winners := []uint{29, 32, 34, 78, 39}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		boolMap.CalculateWinners(&winners, 2)
	}
}

func TestCorrectGetValue(t *testing.T) {
	boolMap := NewBoolMap(1, 90, 100)

	boolMap.SetValue(90, 3, true)
	boolMap.SetValue(78, 3, true)

	assertEqualBool(t, true, boolMap.GetValue(90, 3), "Incorrect value returned")
	assertEqualBool(t, true, boolMap.GetValue(78, 3), "Incorrect value returned")
	assertEqualBool(t, false, boolMap.GetValue(1, 3), "Incorrect value returned")
}

func TestCorrectGetTotalRecords(t *testing.T) {
	boolMap := NewBoolMap(1, 90, 100)

	for i := 0; i < 10; i++ {
		boolMap.IncrementTotalRecords()
	}

	assertEqualInt(t, 10, int(boolMap.GetTotalRecords()), "Incorrect Total Records returned")
}

func TestBoolMap_CalculateWinners(t *testing.T) {
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
	answerMap := boolMap.CalculateWinners(&winningPicks, 2)

	assertEqualMap(t, expectedResult, answerMap, "Wrong query result")
}

package main

import (
	"testing"
)

func BenchmarkLotteryBetsVisitor_Visit(b *testing.B) {
	bitmap := NewBitMap(1, 90, b.N)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitmap, " ")

	bet := "29 32 34 78 39"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lotteryBetsVisitor.Visit(bet)
	}
}

func TestLotteryBetsVisitor_Visit(t *testing.T) {
	bitMap := NewBitMap(1, 90, 100)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitMap, " ")

	mockLotteryBets := []string{
		"29 30 31 33 90",
		"19 60 61 23 89",
	}

	for _, v := range mockLotteryBets {
		lotteryBetsVisitor.Visit(v)
	}

	assertEqualInt(t, 2, int(bitMap.GetTotalRecords()), "Total should be 2")
	// Test bet 1
	assertEqualBool(t, true, bitMap.GetValue(29, 0), "Value should be true")
	assertEqualBool(t, true, bitMap.GetValue(30, 0), "Value should be true")
	assertEqualBool(t, true, bitMap.GetValue(31, 0), "Value should be true")
	assertEqualBool(t, true, bitMap.GetValue(33, 0), "Value should be true")
	assertEqualBool(t, true, bitMap.GetValue(90, 0), "Value should be true")
	// Test bet 2
	assertEqualBool(t, true, bitMap.GetValue(19, 1), "Value should be true")
	assertEqualBool(t, true, bitMap.GetValue(60, 1), "Value should be true")
	assertEqualBool(t, true, bitMap.GetValue(61, 1), "Value should be true")
	assertEqualBool(t, true, bitMap.GetValue(23, 1), "Value should be true")
	assertEqualBool(t, true, bitMap.GetValue(89, 1), "Value should be true")
	// Falsy tests
	assertEqualBool(t, false, bitMap.GetValue(18, 1), "Value should be true")
	assertEqualBool(t, false, bitMap.GetValue(28, 0), "Value should be true")
}

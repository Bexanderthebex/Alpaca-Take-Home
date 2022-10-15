package main

import "testing"

func BenchmarkOptimizedAggregation_Aggregate_10m(b *testing.B) {
	bitmap := NewBitMap(1, 90, b.N)
	lotteryBetsVisitor := NewLotteryBetsVisitor(bitmap, " ")

	for i := 0; i < maximumBettors; i++ {
		lotteryBetsVisitor.Visit("29 10 11 12 13")
	}

	winningPicks := make(map[uint]uint)
	winningPicks[29] = 1
	winningPicks[30] = 1
	winningPicks[32] = 1
	winningPicks[31] = 1
	winningPicks[78] = 1
	agg := NewOptimizedAggregation(bitmap, winningPicks, 2, 5)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		agg.Aggregate()
	}
}

package main

import (
	"os"
	"testing"
)

/*
	variables used in lottery bets engine tests
*/
var mockBets []string
var mockPicks map[uint]uint

func TestMain(m *testing.M) {
	mockBets = []string{
		"29 10 11 12 13",
		"20 31 32 81 60",
		"78 30 29 10 11",
		"29 30 32 31 80",
		"29 30 32 31 78",
	}
	picks := make(map[uint]uint)
	picks[29] = 1
	picks[30] = 1
	picks[32] = 1
	picks[31] = 1
	picks[78] = 1
	mockPicks = picks

	testStatus := m.Run()

	os.Exit(testStatus)
}

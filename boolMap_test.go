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

func TestCorrectGetValue_BoolMap(t *testing.T) {
	boolMap := NewBoolMap(1, 90, 100)

	boolMap.SetValue(90, 3, true)
	boolMap.SetValue(78, 3, true)

	assertEqualBool(t, true, boolMap.GetValue(90, 3), "Incorrect value returned")
	assertEqualBool(t, true, boolMap.GetValue(78, 3), "Incorrect value returned")
	assertEqualBool(t, false, boolMap.GetValue(1, 3), "Incorrect value returned")
}

func TestCorrectGetTotalRecords_BoolMap(t *testing.T) {
	boolMap := NewBoolMap(1, 90, 100)

	for i := 0; i < 10; i++ {
		boolMap.IncrementTotalRecords()
	}

	assertEqualInt(t, 10, int(boolMap.GetTotalRecords()), "Incorrect Total Records returned")
}

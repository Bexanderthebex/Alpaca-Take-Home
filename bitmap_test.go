package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkBitMap_SetValue(b *testing.B) {
	boolMap := NewBitMap(1, 90, b.N)

	for i := 0; i < b.N; i++ {
		boolMap.SetValue(uint(90), uint(i), true)
	}
}

func BenchmarkBitMap_GetValue(b *testing.B) {
	boolMap := NewBitMap(1, 90, b.N)

	for i := 0; i < b.N; i++ {
		boolMap.GetValue(uint(90), uint(i))
	}
}

func TestCorrectGetValue_BitMap(t *testing.T) {
	bitMap := NewBitMap(1, 90, 100)

	bitMap.SetValue(90, 3, true)
	bitMap.SetValue(78, 3, true)

	assert.Equal(t, true, bitMap.GetValue(90, 3), "Incorrect value returned")
	assert.Equal(t, true, bitMap.GetValue(78, 3), "Incorrect value returned")
	assert.Equal(t, false, bitMap.GetValue(1, 3), "Incorrect value returned")
}

func TestCorrectGetTotalRecords_BitMap(t *testing.T) {
	bitMap := NewBitMap(1, 90, 100)

	for i := 0; i < 10; i++ {
		bitMap.IncrementTotalRecords()
	}

	assert.Equal(t, 10, int(bitMap.GetTotalRecords()), "Incorrect Total Records returned")
}

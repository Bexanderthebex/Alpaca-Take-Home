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

package main

import (
	"fmt"
	"testing"
)

func assertEqualBool(t *testing.T, expected bool, actual bool, message string) {
	if expected == actual {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", expected, actual)
	}
	t.Fatal(message)
}

func assertEqualInt(t *testing.T, expected int, actual int, message string) {
	if expected == actual {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", expected, actual)
	}
	t.Fatal(message)
}

func assertEqualMap(t *testing.T, expected map[uint]uint, actual map[uint]uint, message string) {
	for key := range expected {
		if expected[key] != actual[key] {
			t.Fatal(message)
		}
	}

	return
}

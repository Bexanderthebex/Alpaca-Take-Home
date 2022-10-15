package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPicksValidPicksValid(t *testing.T) {
	picks := make(map[uint]uint)

	picks[29] = 1
	picks[32] = 1
	picks[34] = 1
	picks[78] = 1
	picks[39] = 1

	isValid := picksValid(picks)

	assert.Equal(t, true, isValid, "Picks were expected to be valid")
}

func TestPicksValidInsufficientLen(t *testing.T) {
	picks := make(map[uint]uint)

	picks[29] = 1
	picks[32] = 1
	picks[34] = 1
	picks[78] = 1

	isValid := picksValid(picks)

	assert.Equal(t, false, isValid, "Picks were expected to be invalid")
}

func TestPicksValidOccurrenceIsNotUnique(t *testing.T) {
	picks := make(map[uint]uint)

	picks[29] = 1
	picks[32] = 1
	picks[34] = 1
	picks[78] = 1
	picks[28] = 2

	isValid := picksValid(picks)

	assert.Equal(t, false, isValid, "Picks were expected to be invalid")
}

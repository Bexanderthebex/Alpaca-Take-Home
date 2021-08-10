package main

import "testing"

func TestPicksValidPicksValid(t *testing.T) {
	picks := []uint{29, 32, 34, 78, 39}

	isValid := picksValid(picks)

	assertEqualBool(t, true, isValid, "Picks were expected to be valid")
}

func TestPicksValidInsufficientLen(t *testing.T) {
	picks := []uint{29, 32, 34, 78}

	isValid := picksValid(picks)

	assertEqualBool(t, false, isValid, "Picks were expected to be invalid")
}

func TestPicksValidOccurrenceIsNotUnique(t *testing.T) {
	picks := []uint{29, 32, 34, 78, 29}

	isValid := picksValid(picks)

	assertEqualBool(t, false, isValid, "Picks were expected to be invalid")
}

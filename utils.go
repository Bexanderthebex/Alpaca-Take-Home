package main

func picksValid(picks map[uint]uint) bool {
	if len(picks) < lottoPickLength {
		return false
	}

	for pick, occurrence := range picks {
		if !isNumberValid(pick) {
			return false
		}

		if occurrence > 1 {
			return false
		}
	}

	return true
}

func isNumberValid(val uint) bool {
	if val >= minimumValidPick && val <= maximumValidPick {
		return true
	}

	return false
}

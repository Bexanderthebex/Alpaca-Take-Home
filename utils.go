package main

func picksValid(picks []uint) bool {
	if len(picks) < lottoPickLength {
		return false
	}

	picksMap := make(map[uint]int)

	for _, v := range picks {
		if !isNumberValid(int(v)) {
			return false
		}

		picksMap[v] += 1
	}

	for _, v := range picksMap {
		if v > 1 {
			return false
		}
	}

	return true
}

func isNumberValid(val int) bool {
	if val >= int(minimumValidPick) && val <= int(maximumValidPick) {
		return true
	}

	return false
}

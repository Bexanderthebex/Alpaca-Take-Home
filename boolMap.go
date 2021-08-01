package main

type BoolMap struct {
	totalRecords uint
	boolMap      map[uint]*[]bool
}

func NewBoolMap(startingCardinality uint, cardinality uint, setSize int) *BoolMap {
	bm := &BoolMap{
		totalRecords: 0,
		boolMap:      make(map[uint]*[]bool),
	}
	for i := startingCardinality; i <= cardinality; i++ {
		boolArr := make([]bool, setSize)
		bm.setBoolSet(i, &boolArr)
	}

	return bm
}

func (bm *BoolMap) SetValue(index uint, recordId uint, value bool) {
	boolSet := bm.boolMap[index]
	(*boolSet)[recordId] = value
}

func (bm *BoolMap) GetValue(index uint, recordId uint) bool {
	boolSet := bm.boolMap[index]
	return (*boolSet)[recordId]
}

func (bm *BoolMap) GetTotalRecords() uint {
	return bm.totalRecords
}

func (bm *BoolMap) IncrementTotalRecords() uint {
	bm.totalRecords += 1

	return bm.totalRecords
}

func (bm *BoolMap) setBoolSet(index uint, boolSet *[]bool) {
	(*bm).boolMap[index] = boolSet
}

func (bm *BoolMap) CalculateWinners(nums *[]uint, minMatch uint) map[uint]uint {
	winningPicksIndex := make(map[uint]*[]bool)
	// winning picks
	for _, v := range *nums {
		winningPicksIndex[v] = (*bm).boolMap[v]
	}

	groupedWinnersCount := make(map[uint]uint)
	for i := uint(len(*nums)); i >= 2; i-- {
		groupedWinnersCount[i] = 0
	}

	for i := uint(0); i < (*bm).totalRecords; i++ {
		var noOfHits uint = 0
		for _, index := range winningPicksIndex {
			if (*index)[i] {
				noOfHits += 1
			}
		}

		if noOfHits >= minMatch {
			groupedWinnersCount[noOfHits] += 1
		}
	}

	return groupedWinnersCount
}

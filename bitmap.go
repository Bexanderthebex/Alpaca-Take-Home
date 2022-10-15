package main

import "github.com/bits-and-blooms/bitset"

type BitMap struct {
	totalRecords uint
	bitMap       map[uint]*bitset.BitSet
}

func NewBitMap(startingCardinality uint, cardinality uint, setSize int) *BitMap {
	bm := &BitMap{
		totalRecords: 0,
		bitMap:       make(map[uint]*bitset.BitSet),
	}
	for i := startingCardinality; i <= cardinality; i++ {
		bm.bitMap[i] = bitset.New(uint(setSize))
	}

	return bm
}

func (bm *BitMap) SetValue(index uint, recordId uint, value bool) {
	(bm.bitMap[index]).SetTo(recordId, value)
}

func (bm BitMap) GetValue(index uint, recordId uint) bool {
	return (bm.bitMap[index]).Test(recordId)
}

func (bm *BitMap) IncrementTotalRecords() uint {
	bm.totalRecords += 1

	return bm.totalRecords
}

func (bm BitMap) GetTotalRecords() uint {
	return bm.totalRecords
}

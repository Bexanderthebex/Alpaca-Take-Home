package lib

type BitMap map[uint]*[]byte

func New(startingCardinality uint, cardinality uint, setSize int) *BitMap {
	b := make(BitMap)
	for i := startingCardinality; i <= cardinality; i++ {
		byteSet := make([]byte, (setSize/8)+1)
		b[i] = &byteSet
	}

	return &b
}

func (b *BitMap) BitWiseOr(index uint, recordId uint) {
	byteSet := (*b)[index]
	(*byteSet)[recordId/8] |= 128 >> (recordId % 8)
}

func (b *BitMap) GetIndex(index uint) *[]byte {
	return (*b)[index]
}

func (b *BitMap) GetIndexAsBool(index uint) *[]bool {
	byteIndex := (*b)[index]
	boolIndex := make([]bool, len(*byteIndex)*8)

	for indexOfByte, byteValue := range *byteIndex {
		for bit := 0; bit < 8; bit++ {
			if (byteValue<<uint(bit))&128 == 128 {
				boolIndex[8*indexOfByte+bit] = true
			}
		}
	}

	return &boolIndex
}

package main

type CountAggregation struct {
	columns     *map[uint]uint
	accumulator *[]uint
	table       BitMapIndex
}

func NewCountAggregation(winningPicks *map[uint]uint, accumulator *[]uint, table BitMapIndex) *CountAggregation {
	return &CountAggregation{
		columns:     winningPicks,
		accumulator: accumulator,
		table:       table,
	}
}

func (ca *CountAggregation) Aggregate() interface{} {
	for winningPick := range *ca.columns {
		for i := uint(0); i < ca.table.GetTotalRecords(); i++ {
			isTrue := ca.table.GetValue(winningPick, i)
			if isTrue {
				(*ca.accumulator)[i] += 1
			}
		}
	}

	return ca.accumulator
}

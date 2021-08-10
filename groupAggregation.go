package main

type GroupAggregation struct {
	accumulator *[]uint
	groups      []uint
}

func NewGroupAggregation(accumulator *[]uint, groups []uint) *GroupAggregation {
	return &GroupAggregation{
		accumulator: accumulator,
		groups:      groups,
	}
}

func (ga GroupAggregation) Aggregate() interface{} {
	output := make(map[uint]uint)
	if len(ga.groups) <= 0 {
		return output
	}

	for _, group := range ga.groups {
		output[group] = 0
	}

	for _, group := range *ga.accumulator {
		if _, exists := output[group]; exists {
			output[group] += 1
		}
	}

	return output
}

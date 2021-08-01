package main

type AggregationCommand int

const (
	GROUP AggregationCommand = iota + 1
)

func (qa AggregationCommand) String() string {
	return [...]string{"GROUP"}[qa]
}

func (qa AggregationCommand) Int() int {
	return int(qa)
}

type Aggregation interface {
	Aggregate(interface{}, interface{}) *[]uint
}

type BoolQueryAggregation struct {
	aggregationArray   *[]uint
	aggregationCommand AggregationCommand
}

func NewQueryAggregationBool(recordLength uint) *BoolQueryAggregation {
	resultingAggregation := make([]uint, recordLength)

	return &BoolQueryAggregation{
		aggregationArray:   &resultingAggregation,
		aggregationCommand: GROUP,
	}
}

func (qa BoolQueryAggregation) Aggregate(records interface{}, category interface{}) *[]uint {
	if qa.aggregationCommand.Int() == 1 {
		for recordId, record := range *(records).(*[]bool) {
			if record == category.(bool) {
				(*qa.aggregationArray)[recordId] += 1
			}
		}
	}

	return qa.aggregationArray
}

type QueryPlan struct {
	columnsToSelect *[]uint
	aggregationCmd  Aggregation
	minValue        uint
	maxValue        uint
	category        interface{}
}

func (qp QueryPlan) SetColumnsToSelect(newColumns *[]uint) {
	qp.columnsToSelect = newColumns
}

func (qp QueryPlan) SelectGroupStrategy(table *BoolMap) map[uint]uint {
	groupedSelectValues := make(map[uint]uint)

	var aggregatedValues *[]uint
	for _, v := range *qp.columnsToSelect {
		aggregatedValues = qp.aggregationCmd.Aggregate((*table).GetIndex(v), qp.category)
	}

	for _, v := range *aggregatedValues {
		if v >= qp.minValue && v <= qp.maxValue {
			groupedSelectValues[v] += 1
		}
	}

	return groupedSelectValues
}

type LotteryBetsQueryEngine struct {
	boolMap *BoolMap
}

func (l *LotteryBetsQueryEngine) ExecuteQuery(qp QueryPlan) map[uint]uint {
	return qp.SelectGroupStrategy(l.boolMap)
}

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
	Aggregate(BitMapIndex, interface{}, uint) *[]uint
	SetCurrentRecord(uint)
}

type BoolQueryAggregation struct {
	aggregationArray   *[]uint
	aggregationCommand AggregationCommand
	currentRecord      uint
}

func NewQueryAggregationBool(recordLength uint) *BoolQueryAggregation {
	resultingAggregation := make([]uint, recordLength)

	return &BoolQueryAggregation{
		aggregationArray:   &resultingAggregation,
		aggregationCommand: GROUP,
	}
}

func (qa *BoolQueryAggregation) SetCurrentRecord(recordIndex uint) {
	qa.currentRecord = recordIndex
}

func (qa BoolQueryAggregation) getCurrentRecord() uint {
	return qa.currentRecord
}

func (qa BoolQueryAggregation) Aggregate(table BitMapIndex, category interface{}, limit uint) *[]uint {
	if qa.aggregationCommand.Int() == 1 {
		for recordId := uint(0); recordId <= limit; recordId++ {
			tableValue := table.GetValue(qa.getCurrentRecord(), recordId)
			if tableValue == category.(bool) {
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

func (qp QueryPlan) SelectGroupStrategy(table BitMapIndex) map[uint]uint {
	groupedSelectValues := make(map[uint]uint)

	var aggregatedValues *[]uint
	for _, v := range *qp.columnsToSelect {
		qp.aggregationCmd.SetCurrentRecord(v)
		totalRecords := table.GetTotalRecords()
		aggregatedValues = qp.aggregationCmd.Aggregate(table, qp.category, totalRecords)
	}

	for _, v := range *aggregatedValues {
		if v >= qp.minValue && v <= qp.maxValue {
			groupedSelectValues[v] += 1
		}
	}

	return groupedSelectValues
}

type LotteryBetsQueryEngine struct {
	boolMap BitMapIndex
}

func (l *LotteryBetsQueryEngine) ExecuteQuery(qp QueryPlan) map[uint]uint {
	return qp.SelectGroupStrategy(l.boolMap)
}

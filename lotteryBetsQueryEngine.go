package main

type Aggregation interface {
	Aggregate(BitMapIndex, interface{}, uint) *[]uint
	SetCurrentRecord(uint)
}

type QueryAggregation struct {
	aggregationArray   *[]uint
	aggregationCommand AggregationType
	currentRecord      uint
}

func NewQueryAggregation(recordLength uint) *QueryAggregation {
	resultingAggregation := make([]uint, recordLength)

	return &QueryAggregation{
		aggregationArray:   &resultingAggregation,
		aggregationCommand: GROUP,
	}
}

func (qa *QueryAggregation) SetCurrentRecord(recordIndex uint) {
	qa.currentRecord = recordIndex
}

func (qa QueryAggregation) getCurrentRecord() uint {
	return qa.currentRecord
}

func (qa QueryAggregation) Aggregate(table BitMapIndex, category interface{}, limit uint) *[]uint {
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

type Query interface {
	Execute() map[uint]uint
}

type QueryPlan struct {
	columnsToSelect *[]uint
	aggregationCmd  Aggregation
	table           BitMapIndex
	queryTpe        QueryType
	minValue        uint
	maxValue        uint
	category        interface{}
}

func NewQueryPlan(queryType QueryType, category interface{}, table BitMapIndex) *QueryPlan {
	return &QueryPlan{
		columnsToSelect: nil,
		aggregationCmd:  nil,
		minValue:        0,
		maxValue:        0,
		queryTpe:        queryType,
		category:        category,
		table:           table,
	}
}

func (qp *QueryPlan) SetColumnsToSelect(newColumns *[]uint) {
	qp.columnsToSelect = newColumns
}
func (qp *QueryPlan) SetMinValue(value uint) {
	qp.minValue = value
}

func (qp *QueryPlan) SetMaxValue(value uint) {
	qp.maxValue = value
}

func (qp *QueryPlan) SetAggregationStrategy(aggregation Aggregation) {
	qp.aggregationCmd = aggregation
}

func (qp *QueryPlan) SelectGroupStrategy(table BitMapIndex) map[uint]uint {
	groupedSelectValues := make(map[uint]uint)

	var aggregatedValues *[]uint
	totalRecords := table.GetTotalRecords()
	for _, v := range *qp.columnsToSelect {
		qp.aggregationCmd.SetCurrentRecord(v)
		// accumulate the values
		aggregatedValues = qp.aggregationCmd.Aggregate(table, qp.category, totalRecords)
	}

	// count the accumulated values
	for _, v := range *aggregatedValues {
		if v >= qp.minValue && v <= qp.maxValue {
			groupedSelectValues[v] += 1
		}
	}

	return groupedSelectValues
}

func (qp QueryPlan) Execute() map[uint]uint {
	if qp.queryTpe.Int() == 1 {
		if qp.aggregationCmd != nil {
			return qp.SelectGroupStrategy(qp.table)
		}
	}

	return make(map[uint]uint)
}

type LotteryBetsQueryEngine struct {
	bitmapIndex BitMapIndex
}

func (l *LotteryBetsQueryEngine) ExecuteQuery(q Query) map[uint]uint {
	return q.Execute()
}

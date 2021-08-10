package main

type Aggregation interface {
	Aggregate() interface{}
}

type Query interface {
	Execute() map[uint]uint
}

type SelectQueryPlan struct {
	columnsToSelect *map[uint]uint
	aggregationCmd  []Aggregation
	table           BitMapIndex
	minValue        uint
	maxValue        uint
	category        interface{}
}

func NewSelectQueryPlan(table BitMapIndex) *SelectQueryPlan {
	return &SelectQueryPlan{
		columnsToSelect: nil,
		aggregationCmd:  make([]Aggregation, 0),
		minValue:        0,
		maxValue:        0,
		table:           table,
	}
}

func (qp *SelectQueryPlan) SetColumnsToSelect(newColumns *map[uint]uint) {
	qp.columnsToSelect = newColumns
}

func (qp *SelectQueryPlan) SetMinValue(value uint) {
	qp.minValue = value
}

func (qp *SelectQueryPlan) SetMaxValue(value uint) {
	qp.maxValue = value
}

func (qp *SelectQueryPlan) AddAggregationStrategy(aggregation Aggregation) {
	qp.aggregationCmd = append(qp.aggregationCmd, aggregation)
}

func (qp SelectQueryPlan) Execute() map[uint]uint {
	var result interface{}
	for _, agg := range qp.aggregationCmd {
		result = agg.Aggregate()
	}

	return result.(map[uint]uint)
}

type LotteryBetsQueryEngine struct {
	bitmapIndex BitMapIndex
}

func (l *LotteryBetsQueryEngine) ExecuteQuery(q Query) map[uint]uint {
	return q.Execute()
}

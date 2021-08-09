package main

type AggregationType int

const (
	GROUP AggregationType = iota + 1
)

func (qa AggregationType) Int() int {
	return int(qa)
}

type QueryType int

const (
	SELECT QueryType = iota + 1
)

func (qt QueryType) Int() int {
	return int(qt)
}

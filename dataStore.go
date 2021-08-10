package main

type BitMapIndex interface {
	GetValue(uint, uint) bool
	GetTotalRecords() uint
	SetValue(uint, uint, bool)
	IncrementTotalRecords() uint
}

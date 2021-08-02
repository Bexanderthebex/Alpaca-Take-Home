package main

type BitMapSet interface {
}

type BitMapIndex interface {
	GetValue(uint, uint) bool
	GetTotalRecords() uint
}

package main

import (
	"sync"
)

type OptimizedAggregation struct {
	bitmap        *BitMap
	columns       map[uint]uint
	groupRangeMin uint
	groupRangeMax uint
}

func NewOptimizedAggregation(bitmap *BitMap, columns map[uint]uint, groupRangeMin uint, groupRangeMax uint) *OptimizedAggregation {
	return &OptimizedAggregation{
		bitmap:        bitmap,
		columns:       columns,
		groupRangeMin: groupRangeMin,
		groupRangeMax: groupRangeMax,
	}
}

func (ga OptimizedAggregation) Aggregate() interface{} {
	totalRecords := ga.bitmap.GetTotalRecords()
	numberOfGoroutines := uint(32)
	winningPicks := make([]uint, 0)
	for pick := range ga.columns {
		winningPicks = append(winningPicks, pick)
	}

	results := make([]<-chan uint, 0)

	var producerWaitGroup sync.WaitGroup
	var workerWaitGroup sync.WaitGroup
	slice := (totalRecords / numberOfGoroutines) + 1
	for i := uint(0); i < numberOfGoroutines; i++ {
		startRange := slice * i
		endRange := slice * (i + 1)
		var outputChannel <-chan uint

		if endRange > totalRecords {
			outputChannel = ga.producer(&producerWaitGroup, startRange, totalRecords, winningPicks)
		} else {
			outputChannel = ga.producer(&producerWaitGroup, startRange, endRange, winningPicks)
		}

		results = append(results, ga.worker(&workerWaitGroup, outputChannel))
	}

	go func() {
		producerWaitGroup.Wait()
		workerWaitGroup.Wait()
	}()

	answerMap := make(map[uint]uint)
	for _, channel := range results {
		for result := range channel {
			answerMap[result] += 1
		}
	}

	return answerMap
}

func (ga OptimizedAggregation) producer(group *sync.WaitGroup, start uint, end uint, winningPicks []uint) <-chan uint {
	group.Add(1)
	channel := make(chan uint, 4096*3)

	go func() {
		defer group.Done()
		defer close(channel)
		for i := start; i < end; i++ {
			total := uint(0)

			if ga.bitmap.GetValue((winningPicks)[0], i) {
				total += 1
			}
			if ga.bitmap.GetValue((winningPicks)[1], i) {
				total += 1
			}
			if ga.bitmap.GetValue((winningPicks)[2], i) {
				total += 1
			}
			if ga.bitmap.GetValue((winningPicks)[3], i) {
				total += 1
			}
			if ga.bitmap.GetValue((winningPicks)[4], i) {
				total += 1
			}

			if total >= 2 {
				channel <- total
			}
		}
	}()

	return channel
}

func (ga OptimizedAggregation) worker(group *sync.WaitGroup, queue <-chan uint) <-chan uint {
	group.Add(1)
	defer group.Done()
	accumulator := make(chan uint, 4096*3)

	go func() {
		defer close(accumulator)
		for value := range queue {
			accumulator <- value
		}
	}()

	return accumulator
}

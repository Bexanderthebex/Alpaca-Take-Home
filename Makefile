build:
	go build ./main.go ./boolMap.go ./lotteryBetsVisitor.go ./lotteryBetsQueryEngine.go

runAlpacaInputFile:
	./main 10m-v2.txt

test:
	go test -v .

test-benchmark:
	go test -bench=. -benchmem

build:
	go build ./main.go ./lotteryBetsVisitor.go ./lotteryBetsQueryEngine.go ./bitMap.go

runAlpacaInputFile:
	make build
	./main 10m-v2.txt

test-correctness:
	APP_ENVIRONMENT=TEST go test -v .

test-benchmark:
	APP_ENVIRONMENT=TEST go test -bench=. -benchmem

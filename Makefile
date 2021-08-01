build:
	go build ./main.go ./boolMap.go ./lotteryBetsVisitor.go ./lotteryBetsQueryEngine.go

runAlpacaInputFile:
	make build
	./main 10m-v2.txt

test:
	go test -v .

test-benchmark:
	APP_ENVIRONMENT=TEST go test -bench=. -benchmem

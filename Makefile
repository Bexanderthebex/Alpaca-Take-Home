build:
	go build ./main.go ./boolMap.go ./lotteryBetsVisitor.go ./lotteryBetsQueryEngine.go ./bitMap.go ./dataStore.go

run-alpaca-input-file:
	make build
	gunzip 10m-v2.txt.gz -f
	./main 10m-v2.txt

test-correctness:
	APP_ENVIRONMENT=TEST go test -v .

test-benchmark:
	APP_ENVIRONMENT=TEST go test -bench=. -benchmem

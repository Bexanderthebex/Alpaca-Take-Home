build:
	go build ./main.go ./boolMap.go ./lotteryBetsVisitor.go ./lotteryBetsQueryEngine.go ./bitMap.go ./dataStore.go

extract-alpaca-input-file:
	gunzip 10m-v2.txt.gz -f

run-alpaca-input-file:
	make build
	./main 10m-v2.txt

test-correctness:
	APP_ENVIRONMENT=TEST go test -v .

test-benchmark:
	APP_ENVIRONMENT=TEST go test -bench=. -benchmem

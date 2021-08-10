build:
	go build ./main.go ./boolMap.go ./lotteryBetsVisitor.go ./lotteryBetsQueryEngine.go

extract-alpaca-input-file:
	gunzip 10m-v2.txt.gz -f

run-alpaca-input-file:
	make build
	./main 10m-v2.txt

test-correctness:
	go test -v .

test-benchmark:
	go test -bench=. -benchmem

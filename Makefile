build:
	go build ./main.go ./boolMap.go ./lotteryBetsVisitor.go

runAlpacaInputFile:
	./main 10m-v2.txt
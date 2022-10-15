build:
	go build .

init:
	go mod download

extract-input-file:
	gunzip -c 10m-v2.txt.gz >10m-v2.txt

run-input-file:
	make extract-input-file
	make build
	./lottery-winner-finder 10m-v2.txt

run:
	make build
	@echo using $(file)
	./lottery-winner-finder $(file)

test-correctness:
	APP_ENVIRONMENT=TEST go test -v .

test-benchmark:
	APP_ENVIRONMENT=TEST go test -bench=. -benchmem

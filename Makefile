build:
	@go build -o bin/bank

run:
	@./bin/bank
test:
	@go test -v ./...

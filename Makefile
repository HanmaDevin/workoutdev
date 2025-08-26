run: build
	@go run .

test:
	@go test ./tests

build:
	@go build -o bin/app .

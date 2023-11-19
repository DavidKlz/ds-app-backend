build:
	@go build -o bin/dsbackend

run: build
	@./bin/dsbackend

test:
	@go test -v ./...
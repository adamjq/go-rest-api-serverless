BINARY_NAME=bin/main
CMD_PATH=cmd/api/main.go

default: clean test build

generate:
	go generate ./...

test: generate
	go test ./...

clean:
	rm -f src/$(BINARY_NAME)

build:
	GOARCH=amd64 GOOS=linux go build -o $(BINARY_NAME) $(CMD_PATH)

run:
	go run ./cmd/api/main.go

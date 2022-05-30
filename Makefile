
CDK_DIR=ops/cdk
BINARY_NAME=dist/api
CMD_PATH=cmd/api/main.go

default: clean test build

generate:
	go generate ./...

test: generate
	go test ./...

clean:
	rm -rf dist

build:
	GOARCH=amd64 GOOS=linux go build -o $(BINARY_NAME) $(CMD_PATH)

run:
	go run ./cmd/api/main.go

########## AWS ##########

deploy: build
	cd $(CDK_DIR) && npx cdk synth && npx cdk deploy

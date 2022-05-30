
CDK_DIR=ops/cdk
BINARY_NAME=dist/api
CMD_PATH=cmd/api/main.go
OPENAPI_SPEC=openapi/api.yaml

default: clean validate test build

bin/swagger-cli:
	@npm install @apidevtools/swagger-cli --prefix bin --no-save
	@ln -sf node_modules/.bin/swagger-cli bin/swagger-cli

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

validate: bin/swagger-cli
	bin/swagger-cli validate $(OPENAPI_SPEC)


########## AWS ##########

deploy: build
	cd $(CDK_DIR) && npx cdk synth && npx cdk deploy

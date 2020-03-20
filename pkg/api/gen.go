// Package api contains the REST interfaces.
package api

//go:generate env GOBIN=$PWD/.bin GO111MODULE=on go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen
//go:generate $PWD/.bin/oapi-codegen -generate types,server -package api -o users.gen.go ../../openapi/api.yaml
//go:generate gofmt -s -w users.gen.go

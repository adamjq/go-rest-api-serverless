package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adamjq/go-rest-api-serverless/internal/config"
	"github.com/adamjq/go-rest-api-serverless/internal/healthz"
	"github.com/adamjq/go-rest-api-serverless/internal/server"
	"github.com/adamjq/go-rest-api-serverless/pkg/api"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	openapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoLambda *echoadapter.EchoLambda
var cfg *config.Config
var e *echo.Echo

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Echo cold start")
	var err error
	cfg, err = config.NewConfig()
	if err != nil || cfg == nil {
		log.Fatal("error loading config")
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		errMsg := fmt.Sprintf("Error loading swagger spec\n: %s", err)
		log.Fatal(errMsg)
	}

	e = echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(openapimiddleware.OapiRequestValidator(swagger))

	e.GET("/healthz", echo.WrapHandler(healthz.Handler()))

	server := server.Server{}
	api.RegisterHandlers(e, server)

	echoLambda = echoadapter.New(e)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	log.Printf("Main called")
	port := fmt.Sprintf(":%s", cfg.Addr)

	if cfg.IsLambdaRuntime() {
		lambda.Start(Handler)
	} else {
		e.Logger.Fatal(e.Start(port))
	}
}

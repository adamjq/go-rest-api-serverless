package main

import (
	"fmt"
	"log"

	"github.com/adamjq/go-rest-api-serverless/internal/config"
	"github.com/adamjq/go-rest-api-serverless/internal/healthz"
	"github.com/adamjq/go-rest-api-serverless/internal/server"
	"github.com/adamjq/go-rest-api-serverless/pkg/api"
	"github.com/apex/gateway"
	openapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error loading config")
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		errMsg := fmt.Sprintf("Error loading swagger spec\n: %s", err)
		log.Fatal(errMsg)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(openapimiddleware.OapiRequestValidator(swagger))

	e.GET("/healthz", echo.WrapHandler(healthz.Handler()))

	server := server.Server{}
	api.RegisterHandlers(e, server)

	port := fmt.Sprintf(":%s", cfg.Addr)

	if cfg.IsLambdaRuntime() {
		log.Fatal(gateway.ListenAndServe(port, nil))
	} else {
		e.Logger.Fatal(e.Start(port))
	}
}

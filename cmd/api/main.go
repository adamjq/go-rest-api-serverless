package main

import (
	"fmt"
	"log"

	"github.com/adamjq/go-rest-api-serverless/internal/config"
	"github.com/adamjq/go-rest-api-serverless/internal/healthz"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error loading config")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthz", echo.WrapHandler(healthz.Handler()))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Addr)))

	// // Routes
	// e.POST("/users", server.CreateUser(ctx))
	// e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	// if cfg.IsLambdaRuntime() {
	// 	log.Fatal(gateway.ListenAndServe("", mux))
	// }

	// // Start server
	// isLambda := os.Getenv("LAMBDA")
	// if isLambda == "TRUE" {
	// 	lambdaAdapter := &LambdaAdapter{Echo: e}
	// 	lambda.Start(lambdaAdapter.Handler)
	// } else {
	// 	e.Logger.Fatal(e.Start(":1323"))
	// }
}

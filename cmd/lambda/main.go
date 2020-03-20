package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	isLambda := os.Getenv("LAMBDA")
	if isLambda == "TRUE" {
		lambdaAdapter := &LambdaAdapter{Echo: e}
		lambda.Start(lambdaAdapter.Handler)
	} else {
		e.Logger.Fatal(e.Start(":1323"))
	}
}

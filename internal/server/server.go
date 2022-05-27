package server

import (
	"net/http"

	"github.com/adamjq/go-rest-api-serverless/pkg/api"
	"github.com/labstack/echo/v4"
)

type Server struct{}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users = map[int]*user{}
	seq   = 1
)

func (s Server) CreateUser(ctx echo.Context) error {

	newUser := new(api.CreateUserInput)
	if err := ctx.Bind(newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// store the user
	u := &user{
		ID:   seq,
		Name: newUser.Name,
	}
	users[u.ID] = u
	seq++

	return ctx.JSON(http.StatusCreated, u)
}

func (s Server) GetUser(ctx echo.Context, userID int) error {
	if _, ok := users[userID]; ok {
		return ctx.JSON(http.StatusOK, users[userID])
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

func (s Server) ListUsers(ctx echo.Context) error {
	allUsers := make([]user, 0, len(users))
	for u := range users {
		allUsers = append(allUsers, *users[u])
	}
	return ctx.JSON(http.StatusOK, allUsers)
}

func (s Server) UpdateUser(ctx echo.Context, userID int) error {
	updateUserInput := new(api.UpdateUserInput)
	if err := ctx.Bind(updateUserInput); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, ok := users[userID]; ok {

		// update the name
		users[userID].Name = updateUserInput.Name
		return ctx.JSON(http.StatusOK, users[userID])
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

func (s Server) DeleteUser(ctx echo.Context, userID int) error {
	if _, ok := users[userID]; ok {
		delete(users, userID)
		return ctx.NoContent(http.StatusNoContent)
	}
	return echo.NewHTTPError(http.StatusNotFound)
}

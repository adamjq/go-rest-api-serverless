package server

import (
	"fmt"

	"github.com/google/uuid"
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
	fmt.Println("CreateUser: called")
	return nil
}

func (s Server) GetUser(ctx echo.Context, id uuid.UUID) error {
	fmt.Println("GetUser: called")
	return nil
}

func (s Server) ListUsers(ctx echo.Context) error {
	fmt.Println("ListUsers: called")
	return nil
}

func (s Server) UpdateUser(ctx echo.Context, id uuid.UUID) error {
	fmt.Println("UpdateUser: called")
	return nil
}

func (s Server) DeleteUser(ctx echo.Context, id uuid.UUID) error {
	fmt.Println("UpdateUser: called")
	return nil
}

// func CreateUser(c echo.Context) error {
// 	u := &user{
// 		ID: seq,
// 	}
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	users[u.ID] = u
// 	seq++
// 	return c.JSON(http.StatusCreated, u)
// }

// func GetUser(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "error parsing id", err)
// 	}

// 	return c.JSON(http.StatusOK, users[id])
// }

// func UpdateUser(c echo.Context) error {
// 	u := new(user)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "error parsing id", err)
// 	}

// 	users[id].Name = u.Name
// 	return c.JSON(http.StatusOK, users[id])
// }

// func DeleteUser(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "error parsing id", err)
// 	}

// 	delete(users, id)
// 	return c.NoContent(http.StatusNoContent)
// }

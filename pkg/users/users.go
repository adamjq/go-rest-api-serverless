package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	users = map[int]*user{}
	seq   = 1
)

func CreateUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echoerr.NewHTTPError(http.StatusInternalServerError, "error parsing id", err)
	}

	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echoerr.NewHTTPError(http.StatusInternalServerError, "error parsing id", err)
	}

	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echoerr.NewHTTPError(http.StatusInternalServerError, "error parsing id", err)
	}

	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

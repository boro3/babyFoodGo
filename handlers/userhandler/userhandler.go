package userhandler

import (
	"babyFood/pkg/auth"
	"babyFood/pkg/functions/password"
	"babyFood/pkg/user"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	u, err := user.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

func GetUsers(c echo.Context) error {
	users, err := user.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JwtCustomClaims)
	token_id := claims.ID
	if token_id != id {
		return c.JSON(http.StatusForbidden, nil)
	}

	count, err := user.DeleteUser(id)
	if err != nil || count != 1 {
		return c.JSON(http.StatusInternalServerError, count)
	}
	return c.JSON(http.StatusOK, count)
}

func Update(c echo.Context) error {
	id := c.Param("id")

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JwtCustomClaims)
	token_id := claims.ID
	if token_id != id {
		return c.JSON(http.StatusForbidden, nil)
	}

	u, err := user.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	u.Password, err = password.HashPassword(u.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := u.UpdateUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

package authhandler

import (
	"babyFood/pkg/auth"
	"babyFood/pkg/password"
	"babyFood/pkg/user"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//Handler function for creating users from request.
func CreateUser(c echo.Context) error {
	u := new(user.User)
	err := c.Bind(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	u.ID = uuid.NewString()
	validate := validator.New()
	err = validate.Struct(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	u.Password, err = password.HashPassword(u.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = u.CreateUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "Created")
}

//Handler function for issuing jwt tokens if the request provieds corect data.
func Login(c echo.Context) error {
	loginData := new(user.User)
	err := c.Bind(loginData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.StructExcept(loginData, "FirstName", "LastName", "Dob", "ID")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	u, err := user.GetUserByEmail(loginData.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	valid := password.CheckPasswordHash(loginData.Password, u.Password)
	if !valid {
		return c.JSON(http.StatusForbidden, "Wrong Password")
	}

	claims := auth.JwtCustomClaims{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		ID:        u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

package recipehandler

import (
	"babyFood/pkg/auth"
	"babyFood/pkg/recipe"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateRecipe(c echo.Context) error {
	recipe := new(recipe.Recipe)
	err := c.Bind(recipe)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong data")
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(auth.JwtCustomClaims)
	token_id := claims.Id
	recipe.Uid = token_id
	err = recipe.CreateRecipe()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "Created")
}

func GetRecipes(c echo.Context) error {
	recipes, err := recipe.GetRecipes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, recipes)
}

func GetNewRecipes(c echo.Context) error {
	recipes, err := recipe.GetNewRecipes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, recipes)
}

func GetRecipe(c echo.Context) error {
	id := c.Param("id")
	r, err := recipe.GetRecipe(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}

func GetUserRecipes(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(auth.JwtCustomClaims)
	token_id := claims.Id

	recipes, err := recipe.GetRecipesByUid(token_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, recipes)
}

func DeleteRecipe(c echo.Context) error {
	id := c.Param("id")

	r, err := recipe.GetRecipe(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "recipe does not exsist")
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(auth.JwtCustomClaims)
	token_id := claims.Id
	if token_id != r.Uid {
		return c.JSON(http.StatusForbidden, nil)
	}

	count, err := recipe.DeleteRecipe(id)
	if err != nil || count != 1 {
		return c.JSON(http.StatusInternalServerError, count)
	}
	return c.JSON(http.StatusOK, count)
}

func UpdateRecipe(c echo.Context) error {
	id := c.Param("id")

	r, err := recipe.GetRecipe(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect id param")
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(auth.JwtCustomClaims)
	token_id := claims.Id
	if token_id != r.Uid {
		return c.JSON(http.StatusForbidden, nil)
	}

	err = c.Bind(&r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong data")
	}

	res, err := r.UpdateRecipe()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

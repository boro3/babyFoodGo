package recipehandler

import (
	"babyFood/pkg/auth"
	"babyFood/pkg/recipe"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateRecipe(c echo.Context) error {
	recipe := new(recipe.Recipe)
	err := c.Bind(recipe)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong data")
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JwtCustomClaims)
	token_id := claims.ID
	recipe.Uid = token_id
	recipe.ID = uuid.NewString()

	validate := validator.New()
	err = validate.Struct(recipe)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = recipe.CreateRecipe()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "Created")
}

//Handler function for getting all recipes from database on request.
func GetRecipes(c echo.Context) error {
	recipes, err := recipe.GetRecipes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, recipes)
}

//Handler function for getting three newest recipes from database on request.
func GetNewRecipes(c echo.Context) error {
	recipes, err := recipe.GetNewRecipes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, recipes)
}

//Handler function for getting single recipe from  database on request id param.
func GetRecipe(c echo.Context) error {
	id := c.Param("id")
	r, err := recipe.GetRecipe(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}

//Handler function for getting all recipes for jwt from database on request.
func GetUserRecipes(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JwtCustomClaims)
	token_id := claims.ID

	recipes, err := recipe.GetRecipesByUid(token_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, recipes)
}

//Handler function for getting all recipes for category queryparam from database on request.
func GetRecipesByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	r, err := recipe.GetRecipesByCategory(category)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, r)
}

//Handler function for deleing recipe from database for provided id param on request
func DeleteRecipe(c echo.Context) error {
	id := c.Param("id")

	r, err := recipe.GetRecipe(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "recipe does not exsist")
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JwtCustomClaims)
	token_id := claims.ID
	if token_id != r.Uid {
		return c.JSON(http.StatusForbidden, nil)
	}

	count, err := recipe.DeleteRecipe(id)
	if err != nil || count != 1 {
		return c.JSON(http.StatusInternalServerError, count)
	}
	return c.JSON(http.StatusOK, count)
}

//Handler function for updating recipe from database for provided id param on request and valid jwt
func UpdateRecipe(c echo.Context) error {
	id := c.Param("id")

	r, err := recipe.GetRecipe(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "incorrect id param")
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JwtCustomClaims)
	token_id := claims.ID
	if token_id != r.Uid {
		return c.JSON(http.StatusForbidden, nil)
	}

	err = c.Bind(&r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong data")
	}

	validate := validator.New()
	err = validate.Struct(r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := r.UpdateRecipe()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

//Handler fucnction for incrementing field "stars" on recipe for provided id param on request.
func IncrementStars(c echo.Context) error {
	id := c.Param("id")

	r, err := recipe.GetRecipe(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Recipe does not exsist")
	}

	res, err := r.IncrementRecipeStars()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

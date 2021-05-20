package main

import (
	"babyFood/handlers/authhandler"
	"babyFood/handlers/recipehandler"
	"babyFood/handlers/storagehandler"
	"babyFood/handlers/userhandler"
	"babyFood/pkg/auth"
	"babyFood/pkg/db"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.InitializeDBConnection()

	e := echo.New()

	u := e.Group("/user")

	config := middleware.JWTConfig{
		Claims:     &auth.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	}

	u.Use(middleware.JWTWithConfig(config))

	u.GET("", userhandler.GetUsers)
	u.GET("/:id", userhandler.GetUser)
	u.DELETE("/:id", userhandler.DeleteUser)
	u.PUT("/:id", userhandler.Update)

	r := e.Group("/recipe")

	r.GET("", recipehandler.GetRecipes)
	r.POST("", recipehandler.CreateRecipe, middleware.JWTWithConfig(config))
	r.PATCH("/:id", recipehandler.IncrementStars)
	r.PUT("/:id", recipehandler.UpdateRecipe, middleware.JWTWithConfig(config))
	r.DELETE("/:id", recipehandler.DeleteRecipe, middleware.JWTWithConfig(config))
	r.GET("/new", recipehandler.GetNewRecipes)
	r.GET("/user", recipehandler.GetUserRecipes, middleware.JWTWithConfig(config))
	r.GET("/:id", recipehandler.GetRecipe)

	e.POST("/create", authhandler.CreateUser)
	e.POST("/login", authhandler.Login)

	e.POST("/upload/profile", storagehandler.UploadProfilePicture)
	e.GET("/download/profile/:img", storagehandler.DownloadProfilePicture)
	e.POST("/upload/recipe", storagehandler.UploadRecipePicture)
	e.GET("/download/recipe/:img", storagehandler.DownloadProfilePicture)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Start(":" + port)
	fmt.Printf("Running...")

}

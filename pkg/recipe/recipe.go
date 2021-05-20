package recipe

import (
	"babyFood/pkg/db"
	"errors"
	"time"
)

type Recipe struct {
	ID          string    `json:"_id" db:"_id" validate:"required,uuid"`
	Uid         string    `json:"uid" db:"uid" validate:"required,uuid"`
	Description string    `json:"description" db:"description" validate:"required"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Recipe      string    `json:"recipe" db:"recipe"`
	Prep_time   int       `json:"prep_time" db:"prep_time" validate:"required"`
	Stars       int       `json:"stars" db:"stars"`
	Persons     int       `json:"persons" db:"persons" validate:"required"`
	Image       *string   `json:"img" db:"img"`
	Created     time.Time `json:"_created" db:"_created"`
	Deleted     bool      `json:"_deleted" db:"_deleted"`
}

//Function for getting single recipe from database.
//String id of the recipe requestet as input
//If recipe found in databse returns Recipe struct or empty recipe struct if recipe is not found
func GetRecipe(id string) (Recipe, error) {
	var recipe Recipe
	err := db.DBClient.Get(&recipe, getRecipeQuery, id)
	if err != nil {
		return recipe, err
	}
	return recipe, nil
}

//Function for getting all recipe from database.
//If successful returns array of Recipe struct if not error
//Recipes are ordered by "stars" in descending order. If there are no recipes empty array is returned
func GetRecipes() ([]Recipe, error) {
	var recipes []Recipe
	err := db.DBClient.Select(&recipes, getRecipesQuery)
	if err != nil {
		return recipes, err
	}
	return recipes, nil
}

//Function for getting three newest recipe from database.
//If successful returns array of Recipe struct if not error
//Recipes are ordered by _created in descending order if no recipes are found empty array is returned
func GetNewRecipes() ([]Recipe, error) {
	var recipes []Recipe
	err := db.DBClient.Select(&recipes, getNewRecipesQuery)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

//Function for writing recipe into the database.
//The fucntion is attached to Recipe structure.
func (r Recipe) CreateRecipe() error {
	_, err := db.DBClient.NamedExec(createRecipeQuery, r)
	if err != nil {
		return err
	}
	return nil
}

//Function for getting all recipe from database of single user.
//String of the user id as input is required
//returns array of recipes or empty array if there is none
func GetRecipesByUid(uid string) ([]Recipe, error) {
	var recipes []Recipe
	err := db.DBClient.Select(&recipes, getRecipesByUid, uid)
	if err != nil {
		return recipes, err
	}
	return recipes, nil
}

//Fucntion for deleting recipe from database.
//String of the recipe id is required as input
//Returns number int64 of the rows affected in the database
func DeleteRecipe(id string) (int64, error) {
	var a bool = true
	res, err := db.DBClient.Exec(deleteRecipeQuery, a, id)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

//Function for updating recipe record in database.
//String of the recipe id is required as input.
//Returns number int64 of the rows affected in the database.
func (r Recipe) UpdateRecipe() (int64, error) {
	res, err := db.DBClient.NamedExec(updateRecipeQuery, r)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

//Function for incrementing field "stars" from recipe record.
//This fucnion is attached to structure Recipe.
//Returns number int64 of the rows affected in the database.
func (r Recipe) IncrementRecipeStars() (int64, error) {
	res, err := db.DBClient.NamedExec(incrementRecipeStarsQuerry, r)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

var getRecipeQuery = "SELECT * FROM recipe WHERE _id = ? AND _deleted = false;"
var getRecipesQuery = "SELECT * FROM recipe WHERE _deleted = false ORDER BY stars DESC;"
var getNewRecipesQuery = "SELECT * FROM recipe WHERE _deleted = false ORDER BY _created DESC LIMIT 3;"
var getRecipesByUid = "SELECT * FROM recipe WHERE uid = ? AND _deleted = false;"

var createRecipeQuery = `INSERT INTO recipe 
 (_id, uid, description , title, recipe , prep_time , persons) 
VALUES(:_id, :uid, :description, :title, :recipe, :prep_time, :persons);`

var deleteRecipeQuery = "UPDATE recipe SET _deleted = ? WHERE _id = ?;"

var updateRecipeQuery = `UPDATE recipe SET 
	description = :description ,
	title = :title,
	recipe = :recipe,
	prep_time = :prep_time ,
	persons = :persons ,
	img = :img
WHERE _id =:_id;`

var incrementRecipeStarsQuerry = `UPDATE recipe SET stars = stars + 1 WHERE _id = :_id`

package recipe

import (
	"babyFood/pkg/db"
	"errors"
	"time"
)

type Recipe struct {
	ID          string    `json:"_id" db:"_id"`
	Uid         string    `json:"uid" db:"uid"`
	Description string    `json:"description" db:"description"`
	Title       string    `json:"title" db:"title"`
	Recipe      string    `json:"recipe" db:"recipe"`
	Prep_time   int       `json:"prep_time" db:"prep_time"`
	Stars       int       `json:"stars" db:"stars"`
	Persons     int       `json:"persons" db:"persons"`
	Image       *string   `json:"img" db:"img"`
	Created     time.Time `json:"_created" db:"_created"`
	Deleted     bool      `json:"_deleted" db:"_deleted"`
}

func GetRecipe(id string) (Recipe, error) {
	var recipe Recipe
	err := db.DBClient.Get(&recipe, getRecipeQuery, id)
	if err != nil {
		return recipe, err
	}
	return recipe, nil
}

func GetRecipes() ([]Recipe, error) {
	var recipes []Recipe
	err := db.DBClient.Select(&recipes, getRecipesQuery)
	if err != nil {
		return recipes, err
	}
	return recipes, nil
}

func GetNewRecipes() ([]Recipe, error) {
	var recipes []Recipe
	err := db.DBClient.Select(&recipes, getNewRecipesQuery)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r Recipe) CreateRecipe() error {
	_, err := db.DBClient.NamedExec(createRecipeQuery, r)
	if err != nil {
		return err
	}
	return nil
}

func GetRecipesByUid(uid string) ([]Recipe, error) {
	var recipes []Recipe
	err := db.DBClient.Select(&recipes, getRecipesByUid, uid)
	if err != nil {
		return recipes, err
	}
	return recipes, nil
}

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

func (r Recipe) UpdateRecipe() (int64, error) {
	res, err := db.DBClient.NamedExec(updateRecipeQuery, r)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, errors.New(err.Error())
	}
	return count, nil
}

var getRecipeQuery = "SELECT * FROM recipe WHERE _id = ?;"
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

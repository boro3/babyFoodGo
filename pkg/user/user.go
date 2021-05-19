package user

import (
	"babyFood/pkg/db"
	"time"
)

type User struct {
	ID        string    `json:"_id" db:"_id"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Password  string    `json:"password" db:"password" validate:"required"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required"`
	Dob       time.Time `json:"dob" db:"dob"`
	Image     *string   `json:"img" db:"img"`
	Created   time.Time `json:"_created" db:"_created"`
	Deleted   bool      `json:"_deleted" db:"_deleted"`
}

func GetUser(id string) (User, error) {
	var user User
	err := db.DBClient.Get(&user, getUserQueryStr, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUsers() ([]User, error) {
	var users []User
	err := db.DBClient.Select(&users, getUsersQueryStr)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (u User) SaveUser() error {
	_, err := db.DBClient.NamedExec(createUserQueryStr, u)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) (int64, error) {
	var a bool = true
	res, err := db.DBClient.Exec(deleteUserQueryStr, a, id)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

func (u User) UpdateUser() (int64, error) {
	res, err := db.DBClient.NamedExec(UpdateUserQueryStr, u)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, err
	}
	return count, nil
}

var getUserQueryStr = "SELECT * FROM users WHERE _id = ?;"

var getUsersQueryStr = "SELECT _id, first_name, last_name, email, dob FROM users WHERE _deleted = false;"

var createUserQueryStr = `INSERT INTO users (_id, email, password , first_name, last_name, dob , img) VALUES(:_id, :email, :password, :first_name, :last_name, :dob, :img);`

var deleteUserQueryStr = "UPDATE users SET _deleted = ? WHERE _id = ?;"

var UpdateUserQueryStr = `UPDATE users SET 
	first_name = :first_name ,
	last_name = :last_name,
	email = :email,
	dob= :dob ,
	img = :img,
	password = :password
WHERE _id =:_id;`

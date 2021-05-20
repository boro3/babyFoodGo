package user

import (
	"babyFood/pkg/db"
	"time"
)

type User struct {
	ID        string    `json:"_id" db:"_id" validate:"required,uuid"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Password  string    `json:"password" db:"password" validate:"required,gte=4"`
	FirstName string    `json:"first_name" db:"first_name" validate:"required"`
	LastName  string    `json:"last_name" db:"last_name" validate:"required"`
	Dob       time.Time `json:"dob" db:"dob" validate:"required"`
	Image     *string   `json:"image" db:"image"`
	Created   time.Time `json:"_created" db:"_created"`
	Deleted   bool      `json:"_deleted" db:"_deleted"`
}

//GetUser from the database for provided id string as input.
//If there is no resul found in the database empty User is returned
func GetUser(id string) (User, error) {
	var user User
	err := db.DBClient.Get(&user, getUserQueryStr, id)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetUsers from the database.
//If there is no resul found in the database empty array is returned.
func GetUsers() ([]User, error) {
	var users []User
	err := db.DBClient.Select(&users, getUsersQueryStr)
	if err != nil {
		return users, err
	}
	return users, nil
}

//GetUser from the database for provided email  string as input.
//If there is no resul found in the database empty User is returned
func GetUserByEmail(email string) (User, error) {
	var user User
	err := db.DBClient.Get(&user, getUserByEmailQueryStr, email)
	if err != nil {
		return user, err
	}
	return user, nil
}

//CreateUser write new user in the database.
//Returns error if unsuccessful.
func (u User) CreateUser() error {
	_, err := db.DBClient.NamedExec(createUserQueryStr, u)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

//DeleteUser from the database. String id as input is required.
//Returns number of rows affected if unsuccessful 0 is returned
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

//UpdateUser in the database. Attached function to User struct.
//Returns count of rows affected as result.
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

var createUserQueryStr = `INSERT INTO users (_id, email, password , first_name, last_name, dob , image) VALUES(:_id, :email, :password, :first_name, :last_name, :dob, :image);`

var deleteUserQueryStr = "UPDATE users SET _deleted = ? WHERE _id = ?;"

var UpdateUserQueryStr = `UPDATE users SET 
	first_name = :first_name ,
	last_name = :last_name,
	email = :email,
	dob= :dob ,
	image = :image,
	password = :password
WHERE _id =:_id;`

var getUserByEmailQueryStr = "SELECT * FROM users WHERE email = ? AND _deleted = false;"

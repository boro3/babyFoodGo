package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitializeDBConnection() {
	databaseUserName := os.Getenv("DB_USERNAME")
	databaseHost := os.Getenv("DB_HOST")
	databasePort := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_NAME")
	databasePassword := os.Getenv("DB_PASSWORD")
	connectionString := (databaseUserName + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?parseTime=true")
	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}

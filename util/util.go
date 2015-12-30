package util

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

// Create a new mongo session and panics if connection error occurs
func GetMongoSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection fails, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}

func GetMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/zencart?charset=utf8")

	if err != nil {
		panic(err)
	}

	return db
}

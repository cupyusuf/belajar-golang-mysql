package main

import (
	"database/sql"
	"time"
)

func GetConnnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}

package belajar_golang_mysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
